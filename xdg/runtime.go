package xdg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jgoguen/go-utils/pathutils"
)

// RuntimeDirVar holds the environment variable with the location of the user
// runtime files directory.
const RuntimeDirVar = "XDG_RUNTIME_DIR"

// RuntimeDir gets the user runtime directory.
// On Windows and macOS, this directory must be owned by the current user and
// must exist. On other platforms (Linux, BSD, etc.) this directory must
// additionally have UNIX permissions 0700. If these conditions are not met,
// the directory will not be used.
// If $XDG_RUNTIME_DIR cannot be used, the first available system temporary
// directory will be chosen. If no temporary directory can be found, an empty
// string is returned and callers should refuse to do any operation depending
// on the runtime directory.
func RuntimeDir() string {
	// env.GetenvDefault() can't be used here because according to the XDG
	// specification an error message must be printed if "XDG_RUNTIME_DIR" is
	// not set or points to a non-existent location.
	runtimeDir := os.Getenv(RuntimeDirVar)

	if runtimeDir != "" {
		info, err := os.Stat(runtimeDir)
		if err == nil && info.Mode() == 0700 {
			runtimeDir, err = pathutils.SanitizePath(runtimeDir)
			if err == nil {
				return runtimeDir
			}
		}
	}

	tempDir, err := pathutils.SanitizePath(os.TempDir())
	if err == nil {
		fmt.Fprintf(
			os.Stderr,
			"XDG_RUNTIME_DIR not set or invalid, falling back to %s",
			tempDir,
		)

		return tempDir
	}

	return ""
}

// FindRuntimePath returns the absolute path of the given path relative to the
// runtime directory it's in. If the path is not under the runtime directory, an
// empty string is returned.
// path may be a file of any kind or a directory
func FindRuntimePath(path string) string {
	// findFile() sanitizes the path, it can be directly returned
	return findFile([]string{RuntimeDir()}, path)
}

// GetRuntimePath returns the absolute path resulting from joining the user
// runtime directory with the given relative path. If one can't be found, an
// empty string is returned.
func GetRuntimePath(path string) string {
	dir := RuntimeDir()
	if dir == "" {
		return ""
	}

	candidatePath, err := pathutils.SanitizePath(filepath.Join(dir, path))
	if err == nil {
		return candidatePath
	}

	return ""
}
