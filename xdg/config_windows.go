// +build windows

package xdg

import (
	"path/filepath"
)

var (
	// ConfigHomeVars is the list of environment variables which could contain
	// the user config directory.
	ConfigHomeVars = []string{ConfigHomeVar, "APPDATA"}

	// DefaultHomedirConfigPath is the default path of the user cache directory
	// relative to the user home directory.
	DefaultHomedirConfigPath = filepath.Join("AppData", "Roaming")

	// DefaultSystemConfigDirString is the default value of the environment
	// variable `ConfigDirsVar` if it's not set.
	DefaultSystemConfigDirString = getProgramDataDir()
)
