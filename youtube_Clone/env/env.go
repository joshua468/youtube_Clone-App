package env

import (
	"os"
)

var (
	DBHost     = getEnv("DB_HOST", "localhost")
	DBPort     = getEnv("DB_PORT", "5432")
	DBUser     = getEnv("DB_USER", "joshua468")
	DBPassword = getEnv("DB_PASSWORD", "Temitope2080")
	DBName     = getEnv("DB_NAME", "youtube_Clone")
)

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
