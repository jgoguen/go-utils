package kvstore

import (
	"sync"

	"github.com/peterbourgon/diskv"
)

// Database defines an interface for a database utility
type Database interface {
	// Load loads the data from the DB with the specified key.
	Load(key string) ([]byte, error)

	// Save puts bytes in the bucket. For a new item, set key to the empty
	// string if you want a key automatically chosen for you. Returns the key
	// used for saving data and any error that occurred.
	Save(key string, value []byte) (string, error)

	// HasKey returns true if the key is in the DB, false if not
	HasKey(key string) bool

	// Delete deletes a key from the database
	Delete(key string) error

	// List returns all keys
	List() ([]string, error)

	// Close disconnects from the database and cleans up
	Close() error
}

type diskKVStore struct {
	sync.RWMutex

	db   *diskv.Diskv
	name string
}
