package services

import (
	"context"
	"errors"

	"github.com/officebite/officebite/apps/api/internal/models"
	"github.com/officebite/officebite/apps/api/internal/repository"
	"github.com/officebite/officebite/apps/api/internal/utils"
	"gorm.io/gorm"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

type AuthService struct {
	users     repository.UserRepository
	jwtSecret string
	jwtIssuer string
}

type LoginResult struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func NewAuthService(users repository.UserRepository, jwtSecret string, jwtIssuer string) AuthService {
	return AuthService{users: users, jwtSecret: jwtSecret, jwtIssuer: jwtIssuer}
}

func (s AuthService) Login(ctx context.Context, email string, password string) (*LoginResult, error) {
	user, err := s.users.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}

		return nil, err
	}

	if !utils.CheckPassword(user.PasswordHash, password) {
		return nil, ErrInvalidCredentials
	}
	if !user.IsActive {
		return nil, ErrInvalidCredentials
	}

	token, err := utils.GenerateToken(s.jwtSecret, s.jwtIssuer, *user)
	if err != nil {
		return nil, err
	}

	return &LoginResult{Token: token, User: *user}, nil
}

func (s AuthService) CurrentUser(ctx context.Context, userID uint) (*models.User, error) {
	return s.users.FindByID(ctx, userID)
}
