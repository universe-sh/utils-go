package env

import (
	"fmt"
	"os"
)

// Check variables
func Check(a []string) error {
	for _, env := range a {
		if os.Getenv(env) == "" {
			return fmt.Errorf("Env variable %s is empty or unknown", env)
		}
	}

	return nil
}

// Get variable
func Get(a, b string) string {
	if value := os.Getenv(a); value != "" {
		return value
	}

	return b
}
