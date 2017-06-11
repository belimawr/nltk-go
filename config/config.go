package config

import (
	"fmt"
	"os"
)

// PostgresConnectionString - Returns a postgres connection string
func PostgresConnectionString() string {
	username := getenv("DB_USERNAME", "postgres")
	password := getenv("DB_PASSWORD", "123mudar")
	dbName := getenv("DB_DATABASE", "floresta")
	sslMode := getenv("DB_SSL_MODE", "disable")
	host := getenv("DB_HOSTNAME", "postgres")

	conn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		username,
		password,
		host,
		dbName,
		sslMode)

	return conn
}

func getenv(env, defaultValue string) string {
	if value := os.Getenv(env); value != "" {
		return value
	}
	return defaultValue
}
