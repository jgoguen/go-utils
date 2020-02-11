// +build windows

package xdg

import (
	"path/filepath"

	"github.com/jgoguen/go-utils/env"
	"github.com/jgoguen/go-utils/pathutils"
)

var (
	// UserEnvVars is the list of environment variables which could contain the
	// name of the user, in order of their preference of use.
	UserEnvVars = []string{"USERNAME", "USER"}

	// HomedirEnvVars is the list of environment variables which could contain
	// the user home directory, in order of their preference of use.
	HomedirEnvVars = []string{"USERPROFILE", "HOMEPATH"}

	// DefaultHomedirDirectory is the default path where user home directories
	// are contained.
	DefaultHomedirDirectory = getHomedirContainer()
)

// Getting the home directory location on Windows isn't necessarily simple. This
// wraps the logic for getting the Windows home directory.
func getHomedirContainer() string {
	// The drive to use should be here. If %HOMEDRIVE% isn't set, assuming C:
	// is generally safe.
	homedrive := env.GetenvDefault("HOMEDRIVE", "C:")
	if homedrive != "" {
		homedir, err := pathutils.SanitizePath(filepath.Join(homedrive, "Users"))
		if err != nil {
			return ""
		}
		return homedir
	}

	return ""
}
