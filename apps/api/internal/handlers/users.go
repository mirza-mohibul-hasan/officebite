package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/officebite/officebite/apps/api/internal/models"
	"github.com/officebite/officebite/apps/api/internal/services"
	"gorm.io/gorm"
)

type UserHandler struct {
	users services.UserService
}

type userRequest struct {
	Name       string          `json:"name" binding:"required"`
	Email      string          `json:"email" binding:"required,email"`
	Password   string          `json:"password"`
	Role       models.UserRole `json:"role" binding:"required"`
	Department string          `json:"department"`
	IsActive   *bool           `json:"is_active"`
}

func NewUserHandler(users services.UserService) UserHandler {
	return UserHandler{users: users}
}

func (h UserHandler) List(c *gin.Context) {
	users, err := h.users.List(c.Request.Context())
	if err != nil {
		respondInternalError(c, "failed to load users")
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h UserHandler) Create(c *gin.Context) {
	input, ok := bindUserInput(c, true)
	if !ok {
		return
	}
	user, err := h.users.Create(c.Request.Context(), input)
	if err != nil {
		handleUserError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (h UserHandler) Update(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}
	input, ok := bindUserInput(c, false)
	if !ok {
		return
	}
	user, err := h.users.Update(c.Request.Context(), id, input)
	if err != nil {
		handleUserError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func bindUserInput(c *gin.Context, requirePassword bool) (services.UserInput, bool) {
	var req userRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, http.StatusBadRequest, "name, email, and role are required")
		return services.UserInput{}, false
	}
	if requirePassword && len(req.Password) < 8 {
		respondError(c, http.StatusBadRequest, "password must be at least 8 characters")
		return services.UserInput{}, false
	}
	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	return services.UserInput{
		Name:       req.Name,
		Email:      req.Email,
		Password:   req.Password,
		Role:       req.Role,
		Department: req.Department,
		IsActive:   isActive,
	}, true
}

func handleUserError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, services.ErrInvalidUser):
		respondError(c, http.StatusBadRequest, "invalid user")
	case errors.Is(err, gorm.ErrRecordNotFound):
		respondError(c, http.StatusNotFound, "user not found")
	default:
		respondInternalError(c, "user operation failed")
	}
}
