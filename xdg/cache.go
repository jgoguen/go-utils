package xdg

import (
	"os"
	"path/filepath"

	"github.com/jgoguen/go-utils/pathutils"
)

const (
	// CacheHomeEnvVar is the environment variable used to override the default
	// cache directory location.
	CacheHomeEnvVar = "XDG_CACHE_HOME"
)

// CacheHome finds the user cache directory. The default location is
// platform-dependent. An empty string is returned if there isn't enough
// information to determine the directory location.
func CacheHome() string {
	// First, try to use the explicit environment variables
	for _, envvar := range CacheHomeVars {
		cacheDir := os.Getenv(envvar)
		if cacheDir != "" {
			cacheDir, err := pathutils.SanitizePath(cacheDir)
			if err == nil {
				return cacheDir
			}
		}
	}

	// Next try to construct the value based on the home directory
	homedir := Homedir()
	if homedir != "" {
		// Homedir() sanitizes the path, and the homedir-relative location is
		// a simple path known to be valid. No need to re-sanitize.
		return filepath.Join(homedir, DefaultHomedirCachePath)
	}

	return ""
}

// FindCachePath returns the absolute path of the given path relative to the
// cache directory it's in. If the path is not under any cache directory,
// an empty string is returned.
// path may be a file of any kind or a directory
func FindCachePath(path string) string {
	// findFile() sanitizes the path, it can just be directly returned
	return findFile([]string{CacheHome()}, path)
}

// GetCachePath returns the absolute path resulting from joining the user
// cache directory with the given relative path. If one can't be constructed,
// an empty string is returned.
func GetCachePath(path string) string {
	dir := CacheHome()
	if dir == "" {
		return ""
	}

	// We explicitly ignore the error here because SanitizePath only returns
	// both a path and an error when the path doesn't exist. But we may be
	// intentionally requesting the path of a file/directory not yet created.
	candidatePath, _ := pathutils.SanitizePath(filepath.Join(dir, path))
	if candidatePath != "" {
		return candidatePath
	}

	return ""
}
