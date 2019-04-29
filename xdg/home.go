package xdg

import (
	"os"
	"path/filepath"

	"github.com/jgoguen/go-utils/pathutils"
)

// Homedir gets the home directory based on platform-specific environment
// variables, or the best guess based on platform-specific environment
// variables for the username. If no home directory can be determines, an empty
// string is returned.
func Homedir() string {
	for _, envvar := range HomedirEnvVars {
		homedir := os.Getenv(envvar)
		if homedir != "" {
			homedir, err := pathutils.SanitizePath(homedir)
			if err == nil {
				return homedir
			}
		}
	}

	// Finally, make an assumption based on the platform-specific default
	// location for user home directories.
	for _, envvar := range UserEnvVars {
		val := os.Getenv(envvar)
		if val != "" {
			homedir, err := pathutils.SanitizePath(
				filepath.Join(DefaultHomedirDirectory, val),
			)
			if err == nil {
				return homedir
			}
		}
	}

	return ""
}
