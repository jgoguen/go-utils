package xdg

import (
	"os"
	"path/filepath"

	"github.com/jgoguen/go-utils/env"
	"github.com/jgoguen/go-utils/pathutils"
)

const (
	// DataHomeVar is the environment variable used to override the default
	// location of the user data directory.
	DataHomeVar = "XDG_DATA_HOME"

	// DataDirsVar is the environment variable used to override the default
	// system data directories.
	DataDirsVar = "XDG_DATA_DIRS"
)

// DataHome finds the user data directory. The default location is
// platform-specific. An empty string is returned if there isn't enough
// information to determine the directory location.
func DataHome() string {
	// First, try to use the specific environment variables.
	for _, envvar := range DataHomeVars {
		dataDir := os.Getenv(envvar)
		if dataDir != "" {
			dataDir, err := pathutils.SanitizePath(dataDir)
			if err == nil {
				return dataDir
			}
		}
	}

	// Next try to construct the value based on the user home directory
	homedir := Homedir()
	if homedir != "" {
		dataDir, err := pathutils.SanitizePath(
			filepath.Join(homedir, DefaultHomedirDataPath),
		)
		if err == nil {
			return dataDir
		}
	}

	return ""
}

// DataDirs finds all system data directories
func DataDirs() []string {
	systemConfigDirs := env.GetenvDefault(
		DataDirsVar,
		DefaultSystemDataDirString,
	)
	dirs := filepath.SplitList(systemConfigDirs)

	var sanitizedDirs []string
	for _, dname := range dirs {
		sanitizedPath, err := pathutils.SanitizePath(dname)
		if err == nil {
			sanitizedDirs = append(sanitizedDirs, sanitizedPath)
		}
	}

	return sanitizedDirs
}

// AllDataDirs gets all config directories in order of preference
func AllDataDirs() []string {
	var dirs []string

	if userDir := DataHome(); userDir != "" {
		dirs = append(dirs, userDir)
	}

	dirs = append(dirs, DataDirs()...)

	return dirs
}

// FindDataPath returns the absolute path of the given path relative to the
// data directory it's in. If the path is not under any data directory, an
// empty string is returned.
// path may be a file of any kind or a directory
func FindDataPath(path string) string {
	// findFile() sanitizes the path, it can be directly returned
	return findFile(AllDataDirs(), path)
}

// GetDataPath returns the absolute path resulting from joining the highest
// priority data directory with the given relative path. If one can't be
// found, an empty string is returned.
func GetDataPath(path string) string {
	dirs := AllDataDirs()
	if len(dirs) < 1 {
		return ""
	}

	// We explicitly ignore the error here because SanitizePath only returns
	// both a path and an error when the path doesn't exist. But we may be
	// intentionally requesting the path of a file/directory not yet created.
	dataPath, _ := pathutils.SanitizePath(filepath.Join(dirs[0], path))
	if dataPath != "" {
		return dataPath
	}

	return ""
}
