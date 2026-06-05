package services

import (
	"context"
	"errors"
	"strings"

	"github.com/officebite/officebite/apps/api/internal/models"
	"github.com/officebite/officebite/apps/api/internal/repository"
	"github.com/officebite/officebite/apps/api/internal/utils"
)

var ErrInvalidUser = errors.New("invalid user")

type UserService struct {
	users repository.UserRepository
}

type UserInput struct {
	Name       string
	Email      string
	Password   string
	Role       models.UserRole
	Department string
	IsActive   bool
}

func NewUserService(users repository.UserRepository) UserService {
	return UserService{users: users}
}

func (s UserService) List(ctx context.Context) ([]models.User, error) {
	return s.users.List(ctx)
}

func (s UserService) Create(ctx context.Context, input UserInput) (*models.User, error) {
	user, err := buildUser(input, true)
	if err != nil {
		return nil, err
	}
	if err := s.users.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserService) Update(ctx context.Context, id uint, input UserInput) (*models.User, error) {
	user, err := s.users.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	next, err := buildUser(input, false)
	if err != nil {
		return nil, err
	}

	user.Name = next.Name
	user.Email = next.Email
	user.Role = next.Role
	user.Department = next.Department
	user.IsActive = next.IsActive
	if strings.TrimSpace(input.Password) != "" {
		if len(strings.TrimSpace(input.Password)) < 8 {
			return nil, ErrInvalidUser
		}
		hash, err := utils.HashPassword(input.Password)
		if err != nil {
			return nil, err
		}
		user.PasswordHash = hash
	}

	if err := s.users.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func buildUser(input UserInput, requirePassword bool) (*models.User, error) {
	name := strings.TrimSpace(input.Name)
	email := strings.TrimSpace(input.Email)
	password := strings.TrimSpace(input.Password)
	role := input.Role
	if role == "" {
		role = models.RoleEmployee
	}
	if name == "" || email == "" || (requirePassword && len(password) < 8) {
		return nil, ErrInvalidUser
	}
	if role != models.RoleEmployee && role != models.RoleAdmin {
		return nil, ErrInvalidUser
	}

	user := &models.User{
		Name:       name,
		Email:      email,
		Role:       role,
		Department: strings.TrimSpace(input.Department),
		IsActive:   input.IsActive,
	}
	if requirePassword {
		hash, err := utils.HashPassword(password)
		if err != nil {
			return nil, err
		}
		user.PasswordHash = hash
	}

	return user, nil
}
