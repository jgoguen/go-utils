// +build windows

package xdg

import (
	"path/filepath"
)

var (
	// CacheHomeVars is the list of environment variables which could contain
	// the user cache directory.
	CacheHomeVars = []string{CacheHomeEnvVar, "LOCALAPPDATA"}

	// DefaultHomedirCachePath is the default path of the user cache directory
	// relative to the user home directory.
	DefaultHomedirCachePath = filepath.Join("AppData", "Local")
)
