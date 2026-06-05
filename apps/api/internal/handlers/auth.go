package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/officebite/officebite/apps/api/internal/middleware"
	"github.com/officebite/officebite/apps/api/internal/services"
)

type AuthHandler struct {
	auth services.AuthService
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func NewAuthHandler(auth services.AuthService) AuthHandler {
	return AuthHandler{auth: auth}
}

func (h AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "valid email and password are required"})
		return
	}

	result, err := h.auth.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "login failed"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h AuthHandler) Me(c *gin.Context) {
	userID, ok := c.Get(middleware.ContextUserID)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing auth context"})
		return
	}

	user, err := h.auth.CurrentUser(c.Request.Context(), userID.(uint))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
