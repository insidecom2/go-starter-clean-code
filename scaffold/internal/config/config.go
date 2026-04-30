package config

import "os"

// Config holds configuration loaded from environment
type Config struct {
	AppPort     string
	DatabaseURL string
	LogLevel    string
}

// Load reads configuration from environment with sane defaults
func Load() Config {
	return Config{
		AppPort:     getEnv("APP_PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
