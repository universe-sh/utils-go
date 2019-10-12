package system

import (
	"fmt"
	"os"
)

// CheckEnv variables
func CheckEnv(envs []string) error {
	for _, env := range envs {
		if os.Getenv(env) == "" {
			return fmt.Errorf("Env variable %s is empty or unknown", env)
		}
	}

	return nil
}

// GetEnv variable
func GetEnv(name, defValue string) string {
	var value = os.Getenv(name)
	if value != "" {
		return value
	}

	return defValue
}
