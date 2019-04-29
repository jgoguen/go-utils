package xdg

import (
	"os"
	"path/filepath"

	"github.com/jgoguen/go-utils/env"
	"github.com/jgoguen/go-utils/pathutils"
)

const (
	// ConfigHomeVar is the name of the environment variable used to override the
	// default location of the user's config directory.
	ConfigHomeVar = "XDG_CONFIG_HOME"

	// ConfigDirsVar is the name of the environment variable used to override the
	// default system configuration directories.
	ConfigDirsVar = "XDG_CONFIG_DIRS"
)

// ConfigHome finds the user configuration directory. The default location is
// platform-dependent. An empty string is returned if there isn't enough
// information to determine the directory location.
func ConfigHome() string {
	// First, try to use the explicit environment variables
	for _, envvar := range ConfigHomeVars {
		confDir := os.Getenv(envvar)
		if confDir != "" {
			confDir, err := pathutils.SanitizePath(confDir)
			if err == nil {
				return confDir
			}
		}
	}

	// Next try to construct the value based on the home directory
	homedir := Homedir()
	if homedir != "" {
		// Homedir() sanitizes the path, and the homedir-relative location is
		// a simple path known to be valid. No need to re-sanitize.
		return filepath.Join(homedir, DefaultHomedirConfigPath)
	}

	return ""
}

// ConfigDirs finds all system configuration directories
func ConfigDirs() []string {
	systemConfigDirs := env.GetenvDefault(
		ConfigDirsVar,
		DefaultSystemConfigDirString,
	)
	dirs := filepath.SplitList(systemConfigDirs)

	var sanitizedDirs []string
	for _, dname := range dirs {
		sanitizedPath, err := pathutils.SanitizePath(dname)
		if err == nil {
			sanitizedDirs = append(sanitizedDirs, sanitizedPath)
		}
	}

	return sanitizedDirs
}

// AllConfigDirs gets all config directories in order of preference
func AllConfigDirs() []string {
	var dirs []string

	if userDir := ConfigHome(); userDir != "" {
		dirs = append(dirs, userDir)
	}
	dirs = append(dirs, ConfigDirs()...)

	// Both ConfigHome() and ConfigDirs() sanitize their paths, so we don't
	// need to do it again here.
	return dirs
}

// FindConfigPath returns the absolute path of the given path relative to the
// config directory it's in. If the path is not under any configuration
// directory, an empty string is returned.
// path may be a file of any kind or a directory
func FindConfigPath(path string) string {
	// findFile() sanitizes the path, it can just be directly returned
	return findFile(AllConfigDirs(), path)
}

// GetConfigPath returns the absolute path resulting from joining the highest
// priority config directory with the given relative path. If one can't be
// constructed, an empty string is returned.
func GetConfigPath(path string) string {
	dirs := AllConfigDirs()
	if len(dirs) < 1 {
		return ""
	}

	fname, err := pathutils.SanitizePath(filepath.Join(dirs[0], path))
	if err == nil {
		return fname
	}

	return ""
}
