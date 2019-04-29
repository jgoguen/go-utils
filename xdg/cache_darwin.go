// +build darwin

package xdg

var (
	// CacheHomeVars is the list of environment variables which could contain
	// the user cache directory.
	CacheHomeVars = []string{CacheHomeEnvVar}

	// DefaultHomedirCachePath is the default path of the user cache directory
	// relative to the user home directory.
	DefaultHomedirCachePath = "Library/Caches"
)
