package config

import "testing"

func TestValidateRejectsUnsafeProductionDefaults(t *testing.T) {
	cfg := Config{
		Env:         "production",
		Port:        "8080",
		DatabaseURL: "postgres://officebite:officebite@localhost:5432/officebite?sslmode=disable",
		JWTSecret:   defaultJWTSecret,
		WebOrigins:  []string{"https://officebite.example"},
		AutoMigrate: true,
		SeedData:    true,
	}

	if err := cfg.Validate(); err == nil {
		t.Fatal("expected production defaults to be rejected")
	}
}

func TestValidateAcceptsDevelopmentConfig(t *testing.T) {
	cfg := Config{
		Env:         "development",
		Port:        "8080",
		DatabaseURL: "postgres://officebite:officebite@localhost:5432/officebite?sslmode=disable",
		JWTSecret:   defaultJWTSecret,
		WebOrigins:  []string{"http://localhost:5173"},
		AutoMigrate: true,
		SeedData:    true,
	}

	if err := cfg.Validate(); err != nil {
		t.Fatalf("expected development config to be valid: %v", err)
	}
}

func TestSplitOriginsTrimsCSVOrigins(t *testing.T) {
	origins := splitOrigins("http://localhost:5173, https://officebite.example ")
	if len(origins) != 2 {
		t.Fatalf("expected 2 origins, got %d", len(origins))
	}
	if origins[0] != "http://localhost:5173" || origins[1] != "https://officebite.example" {
		t.Fatalf("unexpected origins: %#v", origins)
	}
}
