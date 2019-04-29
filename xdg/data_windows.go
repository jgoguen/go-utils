// +build windows

package xdg

import (
	"path/filepath"
)

var (
	// DataHomeVars is the list of environment variables which could contain
	// the user config directory.
	DataHomeVars = []string{DataHomeVar}

	// DefaultHomedirDataPath is the default path of the user cache directory
	// relative to the user home directory.
	DefaultHomedirDataPath = filepath.Join("AppData", "Roaming")

	// DefaultSystemDataDirString is the default value of the environment
	// variable `ConfigDirsVar` if it's not set.
	DefaultSystemDataDirString = getProgramDataDir()
)
