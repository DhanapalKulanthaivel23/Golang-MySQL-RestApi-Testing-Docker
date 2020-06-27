package utils

import (
	"os"
)

func GetEnvVar(name string, fallbak string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}
	return fallbak
}
