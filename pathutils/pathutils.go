package pathutils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// SanitizePath takes a path and makes sure it's in a sanitized form. An error
// is returned if the path doesn't exist; the path is also returned so callers
// expecting the path to not exist can ignore the error. An empty string is
// returned if the path could not be sanitized.
// Note: This is not a guard against directory traversal attacks. This is only
// intended to give you a sane canonical path name.
func SanitizePath(path string) (string, error) {
	if path == "" {
		curpath, err := filepath.Abs(".")
		if err != nil {
			return "", err
		}
		return curpath, nil
	}

	candidatePath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	_, err = os.Stat(candidatePath)
	if err != nil {
		return candidatePath, err
	}

	return candidatePath, nil
}

// SanitizePathUnderBase takes a path, calls SanitizePath, and ensures the
// resulting path is at or below the specified base path. An error is returned
// if the path doesn't exist; the path is also returned so callers expecting
// the path to not exist can ignore the error. An empty string is returned if
// the path could not be sanitized.
func SanitizePathUnderBase(path, basepath string) (string, error) {
	candidatePath, err := SanitizePath(path)
	// err may not be nil at this point, but as long as candidatePath isn't an
	// empty string we carry on. The error gets bubbled up to the caller so they
	// can decide whether to ignore the error and use the path.
	if candidatePath == "" {
		return "", err
	}

	// Windows and macOS ignore case, everything else is case-sensitive
	validPath := false
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
		validPath = strings.HasPrefix(
			strings.ToLower(candidatePath),
			strings.ToLower(basepath),
		)
	} else {
		validPath = strings.HasPrefix(candidatePath, basepath)
	}

	if validPath {
		return candidatePath, err
	}

	return "", fmt.Errorf("path '%s' is not at or under '%s'", candidatePath, basepath)
}
