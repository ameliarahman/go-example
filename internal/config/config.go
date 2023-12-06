package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	UserPostgres, PasswordPostgres, NamePostgres, HostPostgres string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading env")
	}

	UserPostgres = getEnv("POSTGRES_USERNAME", "")
	PasswordPostgres = getEnv("POSTGRES_PASSWORD", "")
	NamePostgres = getEnv("POSTGRES_NAME", "")
	HostPostgres = getEnv("POSTGRES_HOST", "127.0.0.1")
}

// function to get data environment variable
func getEnv(key string, err string) string {
	value, exists := os.LookupEnv(key)

	if exists {
		return value
	}

	return err
}
