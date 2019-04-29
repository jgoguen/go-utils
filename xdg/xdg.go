package xdg

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/jgoguen/go-utils/env"
	"github.com/jgoguen/go-utils/pathutils"
)

// findFile looks for the first occurrence of `path` in the directories given.
// If the path doesn't exist in any directory, an empty string is returned.
func findFile(dirs []string, path string) string {
	for _, d := range dirs {
		filePath := filepath.Join(d, path)
		_, err := os.Stat(filePath)
		if err == nil {
			fname, err := pathutils.SanitizePath(filePath)
			if err == nil {
				return fname
			}
		}
	}

	return ""
}

// On Windows, the default location for the system data directory is held in
// the %PROGRAMDATA% environment variable, and if that's not set the default is
// %SYSTEMDRIVE%\ProgramData, and if %SYSTEMDRIVE% isn't set 'C:' is a safe
// assumption. This wraps the logic of getting that default path. If called on
// non-Windows systems, an empty string is returned.
func getProgramDataDir() string {
	if runtime.GOOS != "windows" {
		return ""
	}

	progDataDir := os.Getenv("PROGRAMDATA")
	if progDataDir != "" {
		progDataDir, err := pathutils.SanitizePath(progDataDir)
		if err == nil {
			return progDataDir
		}
	}

	sysDrive := env.GetenvDefault("SYSTEMDRIVE", "C:")
	if sysDrive != "" {
		confDir, err := pathutils.SanitizePath(
			filepath.Join(sysDrive, "ProgramData"),
		)
		if err == nil {
			return confDir
		}
	}

	return ""
}
