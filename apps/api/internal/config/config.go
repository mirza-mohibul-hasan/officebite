package config

import "os"

type Config struct {
	Env              string
	Port             string
	DatabaseURL      string
	JWTSecret        string
	WebOrigin        string
	AutoMigrate      bool
	DatabaseLogLevel string
}

func Load() Config {
	return Config{
		Env:              getEnv("APP_ENV", "development"),
		Port:             getEnv("API_PORT", "8080"),
		DatabaseURL:      getEnv("DATABASE_URL", "postgres://officebite:officebite@localhost:5432/officebite?sslmode=disable"),
		JWTSecret:        getEnv("JWT_SECRET", "change-me-in-production"),
		WebOrigin:        getEnv("WEB_ORIGIN", "http://localhost:5173"),
		AutoMigrate:      getEnv("AUTO_MIGRATE", "true") == "true",
		DatabaseLogLevel: getEnv("DATABASE_LOG_LEVEL", "warn"),
	}
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
