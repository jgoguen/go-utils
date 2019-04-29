// +build darwin

package xdg

var (
	// DataHomeVars is the list of environment variables which could contain
	// the user config directory.
	DataHomeVars = []string{DataHomeVar}

	// DefaultHomedirConfigPath is the default path of the user cache directory
	// relative to the user home directory.
	DefaultHomedirDataPath = "Library/Application Support"

	// DefaultSystemConfigDirString is the default value of the environment
	// variable `ConfigDirsVar` if it's not set.
	DefaultSystemDataDirString = "/Library/Application Support"
)
