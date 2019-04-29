package env

import "os"

// GetenvDefault wraps os.LookupEnv to return a default value if the environment
// variable isn't defined.
func GetenvDefault(envvar, defaultValue string) string {
	val, defined := os.LookupEnv(envvar)

	if !defined {
		return defaultValue
	}

	return val
}
