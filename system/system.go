package system

import (
	"os"
)

// GetEnv variable
func GetEnv(name, defValue string) string {
	var value = os.Getenv(name)
	if value != "" {
		return value
	}

	return defValue
}
