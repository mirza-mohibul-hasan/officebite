package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/officebite/officebite/apps/api/internal/services"
	"gorm.io/gorm"
)

type MenuHandler struct {
	menus services.MenuService
}

type menuRequest struct {
	Title         string `json:"title" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Price         int64  `json:"price" binding:"required,min=1"`
	AvailableDate string `json:"available_date" binding:"required"`
}

func NewMenuHandler(menus services.MenuService) MenuHandler {
	return MenuHandler{menus: menus}
}

func (h MenuHandler) ListToday(c *gin.Context) {
	dateParam := c.DefaultQuery("date", time.Now().Format("2006-01-02"))
	date, err := parseDate(dateParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date must use YYYY-MM-DD"})
		return
	}

	menus, err := h.menus.ListByDate(c.Request.Context(), date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load menus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"menus": menus})
}

func (h MenuHandler) ListAll(c *gin.Context) {
	menus, err := h.menus.ListAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load menus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"menus": menus})
}

func (h MenuHandler) Create(c *gin.Context) {
	input, ok := bindMenuInput(c)
	if !ok {
		return
	}

	menu, err := h.menus.Create(c.Request.Context(), input)
	if err != nil {
		handleMenuError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"menu": menu})
}

func (h MenuHandler) Update(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	input, ok := bindMenuInput(c)
	if !ok {
		return
	}

	menu, err := h.menus.Update(c.Request.Context(), id, input)
	if err != nil {
		handleMenuError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"menu": menu})
}

func (h MenuHandler) Delete(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	if err := h.menus.Delete(c.Request.Context(), id); err != nil {
		handleMenuError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func bindMenuInput(c *gin.Context) (services.MenuInput, bool) {
	var req menuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title, description, price, and available_date are required"})
		return services.MenuInput{}, false
	}

	date, err := parseDate(req.AvailableDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "available_date must use YYYY-MM-DD"})
		return services.MenuInput{}, false
	}

	return services.MenuInput{
		Title:         req.Title,
		Description:   req.Description,
		Price:         req.Price,
		AvailableDate: date,
	}, true
}

func parseDate(value string) (time.Time, error) {
	return time.Parse("2006-01-02", value)
}

func parseIDParam(c *gin.Context, name string) (uint, bool) {
	value, err := strconv.ParseUint(c.Param(name), 10, 64)
	if err != nil || value == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return 0, false
	}

	return uint(value), true
}

func handleMenuError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, services.ErrInvalidMenu):
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid menu"})
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "menu not found"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "menu operation failed"})
	}
}
