package services

import (
	"context"
	"errors"
	"testing"

	"github.com/officebite/officebite/apps/api/internal/models"
	"github.com/officebite/officebite/apps/api/internal/utils"
	"gorm.io/gorm"
)

type fakeUserRepository struct {
	user *models.User
}

func (r fakeUserRepository) Create(ctx context.Context, user *models.User) error {
	return nil
}

func (r fakeUserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	if r.user == nil || r.user.Email != email {
		return nil, gorm.ErrRecordNotFound
	}
	return r.user, nil
}

func (r fakeUserRepository) FindByID(ctx context.Context, id uint) (*models.User, error) {
	if r.user == nil || r.user.ID != id {
		return nil, gorm.ErrRecordNotFound
	}
	return r.user, nil
}

func TestAuthServiceLogin(t *testing.T) {
	hash, err := utils.HashPassword("password123")
	if err != nil {
		t.Fatalf("hash password: %v", err)
	}
	auth := NewAuthService(fakeUserRepository{user: &models.User{
		ID:           1,
		Name:         "Employee User",
		Email:        "employee@officebite.local",
		PasswordHash: hash,
		Role:         models.RoleEmployee,
	}}, "test-secret-that-is-long-enough-32", "officebite-test")

	result, err := auth.Login(context.Background(), "employee@officebite.local", "password123")
	if err != nil {
		t.Fatalf("login: %v", err)
	}
	if result.Token == "" {
		t.Fatal("expected token")
	}
	if result.User.PasswordHash != hash {
		t.Fatal("expected service to return the authenticated user")
	}
}

func TestAuthServiceLoginRejectsInvalidPassword(t *testing.T) {
	hash, err := utils.HashPassword("password123")
	if err != nil {
		t.Fatalf("hash password: %v", err)
	}
	auth := NewAuthService(fakeUserRepository{user: &models.User{
		ID:           1,
		Email:        "employee@officebite.local",
		PasswordHash: hash,
		Role:         models.RoleEmployee,
	}}, "test-secret-that-is-long-enough-32", "officebite-test")

	_, err = auth.Login(context.Background(), "employee@officebite.local", "bad-password")
	if !errors.Is(err, ErrInvalidCredentials) {
		t.Fatalf("expected invalid credentials, got %v", err)
	}
}
