package config

import "os"

type Config struct {
	Env               string
	Port              string
	DatabaseURL       string
	JWTSecret         string
	JWTIssuer         string
	WebOrigin         string
	AutoMigrate       bool
	DatabaseLogLevel  string
	SeedData          bool
	SeedAdminEmail    string
	SeedAdminPass     string
	SeedEmployeeEmail string
	SeedEmployeePass  string
}

func Load() Config {
	return Config{
		Env:               getEnv("APP_ENV", "development"),
		Port:              getEnv("API_PORT", "8080"),
		DatabaseURL:       getEnv("DATABASE_URL", "postgres://officebite:officebite@localhost:5432/officebite?sslmode=disable"),
		JWTSecret:         getEnv("JWT_SECRET", "change-me-in-production"),
		JWTIssuer:         getEnv("JWT_ISSUER", "officebite-api"),
		WebOrigin:         getEnv("WEB_ORIGIN", "http://localhost:5173"),
		AutoMigrate:       getEnv("AUTO_MIGRATE", "true") == "true",
		DatabaseLogLevel:  getEnv("DATABASE_LOG_LEVEL", "warn"),
		SeedData:          getEnv("SEED_DATA", "true") == "true",
		SeedAdminEmail:    getEnv("SEED_ADMIN_EMAIL", "admin@officebite.local"),
		SeedAdminPass:     getEnv("SEED_ADMIN_PASSWORD", "password123"),
		SeedEmployeeEmail: getEnv("SEED_EMPLOYEE_EMAIL", "employee@officebite.local"),
		SeedEmployeePass:  getEnv("SEED_EMPLOYEE_PASSWORD", "password123"),
	}
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
