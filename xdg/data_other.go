// +build !windows !darwin

package xdg

var (
	// DataHomeVars is the list of environment variables which could contain
	// the user config directory.
	DataHomeVars = []string{DataHomeVar}

	// DefaultHomedirDataPath is the default path of the user cache directory
	// relative to the user home directory.
	DefaultHomedirDataPath = ".local/share"

	// DefaultSystemDataDirString is the default value of the environment
	// variable `ConfigDirsVar` if it's not set.
	DefaultSystemDataDirString = "/usr/local/share:/usr/share"
)
