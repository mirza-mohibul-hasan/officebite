package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const defaultJWTSecret = "dev-only-change-this-jwt-secret-32"

type Config struct {
	Env               string
	Port              string
	DatabaseURL       string
	JWTSecret         string
	JWTIssuer         string
	WebOrigin         string
	WebOrigins        []string
	AutoMigrate       bool
	DatabaseLogLevel  string
	SeedData          bool
	SeedAdminEmail    string
	SeedAdminPass     string
	SeedEmployeeEmail string
	SeedEmployeePass  string
}

func Load() Config {
	webOrigin := getEnv("WEB_ORIGIN", "http://localhost:5173")

	return Config{
		Env:               getEnv("APP_ENV", "development"),
		Port:              getEnv("API_PORT", "8080"),
		DatabaseURL:       getEnv("DATABASE_URL", "postgres://officebite:officebite@localhost:5432/officebite?sslmode=disable"),
		JWTSecret:         getEnv("JWT_SECRET", defaultJWTSecret),
		JWTIssuer:         getEnv("JWT_ISSUER", "officebite-api"),
		WebOrigin:         webOrigin,
		WebOrigins:        splitOrigins(webOrigin),
		AutoMigrate:       getEnv("AUTO_MIGRATE", "true") == "true",
		DatabaseLogLevel:  getEnv("DATABASE_LOG_LEVEL", "warn"),
		SeedData:          getEnv("SEED_DATA", "true") == "true",
		SeedAdminEmail:    getEnv("SEED_ADMIN_EMAIL", "admin@officebite.local"),
		SeedAdminPass:     getEnv("SEED_ADMIN_PASSWORD", "password123"),
		SeedEmployeeEmail: getEnv("SEED_EMPLOYEE_EMAIL", "employee@officebite.local"),
		SeedEmployeePass:  getEnv("SEED_EMPLOYEE_PASSWORD", "password123"),
	}
}

func (c Config) Validate() error {
	var problems []string

	if c.Env != "development" && c.Env != "test" && c.Env != "production" {
		problems = append(problems, "APP_ENV must be development, test, or production")
	}
	if _, err := strconv.Atoi(c.Port); err != nil {
		problems = append(problems, "API_PORT must be numeric")
	}
	if c.DatabaseURL == "" {
		problems = append(problems, "DATABASE_URL is required")
	}
	if len(c.WebOrigins) == 0 {
		problems = append(problems, "WEB_ORIGIN must include at least one origin")
	}
	for _, origin := range c.WebOrigins {
		parsed, err := url.Parse(origin)
		if err != nil || parsed.Scheme == "" || parsed.Host == "" {
			problems = append(problems, fmt.Sprintf("WEB_ORIGIN contains invalid origin %q", origin))
		}
	}
	if len(c.JWTSecret) < 32 {
		problems = append(problems, "JWT_SECRET must be at least 32 characters")
	}
	if c.Env == "production" {
		if c.JWTSecret == defaultJWTSecret {
			problems = append(problems, "JWT_SECRET must be changed in production")
		}
		if c.AutoMigrate {
			problems = append(problems, "AUTO_MIGRATE must be false in production")
		}
		if c.SeedData {
			problems = append(problems, "SEED_DATA must be false in production")
		}
	}

	if len(problems) > 0 {
		return errors.New(strings.Join(problems, "; "))
	}

	return nil
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func splitOrigins(value string) []string {
	parts := strings.Split(value, ",")
	origins := make([]string, 0, len(parts))
	for _, part := range parts {
		origin := strings.TrimSpace(part)
		if origin != "" {
			origins = append(origins, origin)
		}
	}

	return origins
}
