package utils

import (
	"testing"

	"github.com/officebite/officebite/apps/api/internal/models"
)

func TestGenerateAndParseToken(t *testing.T) {
	user := models.User{
		ID:    42,
		Role:  models.RoleAdmin,
		Email: "admin@officebite.local",
		Name:  "Admin User",
	}

	token, err := GenerateToken("test-secret-that-is-long-enough-32", "officebite-test", user)
	if err != nil {
		t.Fatalf("generate token: %v", err)
	}

	claims, err := ParseToken("test-secret-that-is-long-enough-32", token)
	if err != nil {
		t.Fatalf("parse token: %v", err)
	}

	if claims.UserID != user.ID {
		t.Fatalf("expected user id %d, got %d", user.ID, claims.UserID)
	}
	if claims.Role != models.RoleAdmin {
		t.Fatalf("expected admin role, got %s", claims.Role)
	}
}
