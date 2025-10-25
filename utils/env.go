package utils

import "os"

// GetEnv returns the environment variable value for key or fallback if empty.
func GetEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
