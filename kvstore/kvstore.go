package kvstore

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"

	"github.com/jgoguen/go-utils/xdg"
	"github.com/peterbourgon/diskv"
)

var dbMap map[string]Database

// OpenDB opens a Database connection. `name` may be an absolute path or a
// directory name relative to the user XDG data directory
func OpenDB(name string) (Database, error) {
	db, found := dbMap[name]
	if found && db != nil {
		return db, nil
	}

	var dbDir string

	if filepath.IsAbs(name) {
		dbDir = name
	} else {
		dataDir := xdg.GetDataPath(name)
		if dataDir == "" {
			return nil, errors.New(
				"could not determine the XDG location for the DB",
			)
		}

		if err := os.MkdirAll(dataDir, 0700); err != nil {
			return nil, err
		}

		dbDir = dataDir
	}

	db = &diskKVStore{
		db: diskv.New(diskv.Options{
			BasePath:     dbDir,
			CacheSizeMax: 1024 * 1024, // 1MB cache max
			FilePerm:     0600,
			PathPerm:     0700,
		}),
		name: name,
	}

	dbMap[name] = db

	return db, nil
}

func (diskdb *diskKVStore) Load(key string) ([]byte, error) {
	diskdb.RLock()
	defer diskdb.RUnlock()

	return diskdb.db.Read(key)
}

func (diskdb *diskKVStore) Save(key string, value []byte) (string, error) {
	var savedKey string
	if key != "" {
		savedKey = key
	} else {
		sum := sha256.Sum256(value)
		savedKey = hex.EncodeToString(sum[:])
	}

	diskdb.Lock()
	defer diskdb.Unlock()

	err := diskdb.db.Write(savedKey, value)

	return savedKey, err
}

func (diskdb *diskKVStore) HasKey(key string) bool {
	diskdb.RLock()
	defer diskdb.RUnlock()

	return diskdb.db.Has(key)
}

func (diskdb *diskKVStore) Delete(key string) error {
	diskdb.Lock()
	defer diskdb.Unlock()

	return diskdb.db.Erase(key)
}

func (diskdb *diskKVStore) List() ([]string, error) {
	var keys []string

	diskdb.RLock()
	defer diskdb.RUnlock()

	for v := range diskdb.db.Keys(nil) {
		keys = append(keys, v)
	}

	return keys, nil
}

func (diskdb *diskKVStore) Close() error {
	delete(dbMap, diskdb.name)
	diskdb.db = nil

	return nil
}
