// +build !windows !darwin

package xdg

var (
	// ConfigHomeVars is the list of environment variables which could contain
	// the user config directory.
	ConfigHomeVars = []string{ConfigHomeVar}

	// DefaultHomedirConfigPath is the default path of the user cache directory
	// relative to the user home directory.
	DefaultHomedirConfigPath = ".config"

	// DefaultSystemConfigDirString is the default value of the environment
	// variable `ConfigDirsVar` if it's not set.
	DefaultSystemConfigDirString = "/etc/xdg"
)
