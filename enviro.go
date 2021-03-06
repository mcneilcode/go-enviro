package enviro

import (
	"os"
	"strconv"
	"strings"
)

// Get returns the value of an environment variable or
// the provided default when non-existent as a string.
func Get(name string, defaultValue string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}
	return defaultValue
}

// GetInt returns the value of an environment variable as
// an integer value or the provided default integer.
func GetInt(name string, defaultValue int) int {
	value := Get(name, "")
	if value, err := strconv.Atoi(value); err == nil {
		return value
	}
	return defaultValue
}

// GetInt64 returns the value of an environment variable as
// a 64-bit base 10 value or the provided default.
func GetInt64(name string, defaultValue int64) int64 {
	value := Get(name, "")
	if value, err := strconv.ParseInt(value, 10, 64); err == nil {
		return value
	}
	return defaultValue
}

// GetBool returns the value of an environment variable as
// a boolean value or the provided default boolean.
func GetBool(name string, defaultValue bool) bool {
	value := Get(name, "")
	if val, err := strconv.ParseBool(value); err == nil {
		return val
	}
	return defaultValue
}

// GetFloat32 returns the value of an environment variable as
// a float32 value or the provided default when non-existant.
func GetFloat32(name string, defaultValue float32) float32 {
	value := Get(name, "")
	if value, err := strconv.ParseFloat(value, 32); err == nil {
		return float32(value)
	}
	return defaultValue
}

// GetFloat64 returns the value of an environment variable as
// a float64 value or the provided default when non-existant.
func GetFloat64(name string, defaultValue float64) float64 {
	value := Get(name, "")
	if value, err := strconv.ParseFloat(value, 64); err == nil {
		return value
	}
	return defaultValue
}

// GetSlice returns the value of an environment variable as
// a string slice based on the provided separator. It also takes
// a string slice to use as a fallback default.
func GetSlice(name string, sep string, defaultValue []string) []string {
	value := Get(name, "")

	if value == "" {
		return defaultValue
	}
	return strings.Split(value, sep)
}
