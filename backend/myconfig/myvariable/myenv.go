package myvariable

import (
	"fmt"
	"os"
)

var ENV_PORT string
var ENV_DB_PROJECT string

func GetMyENV(EnvKey string, defaultValue ...string) string {
	value := os.Getenv(EnvKey)
	fmt.Printf("[ENV] %s: %s\n", EnvKey, value)
	if value == "" && len(defaultValue) != 0 {
		value = defaultValue[0]
	}
	return value
}

func SetEnv() {
	ENV_PORT = GetMyENV("PORT", "8080")
	ENV_DB_PROJECT = GetMyENV("DB_PROJECT")
}
