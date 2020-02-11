// +build !windows
// +build !darwin

package xdg

var (
	// UserEnvVars is the list of environment variables which could contain the
	// name of the user, in order of their preference of use.
	UserEnvVars = []string{"SUDO_USER", "USER", "LOGNAME"}

	// HomedirEnvVars is the list of environment variables which could contain
	// the user home directory, in order of their preference of use.
	HomedirEnvVars = []string{"HOME"}

	// DefaultHomedirDirectory is the default path where user home directories
	// are contained.
	DefaultHomedirDirectory = "/home"
)
