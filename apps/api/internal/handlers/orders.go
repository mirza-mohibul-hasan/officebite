package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/officebite/officebite/apps/api/internal/middleware"
	"github.com/officebite/officebite/apps/api/internal/services"
	"gorm.io/gorm"
)

type OrderHandler struct {
	orders services.OrderService
}

type placeOrderRequest struct {
	MenuID uint `json:"menu_id" binding:"required,min=1"`
}

func NewOrderHandler(orders services.OrderService) OrderHandler {
	return OrderHandler{orders: orders}
}

func (h OrderHandler) Place(c *gin.Context) {
	userID, ok := currentUserID(c)
	if !ok {
		return
	}

	var req placeOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "menu_id is required"})
		return
	}

	order, err := h.orders.Place(c.Request.Context(), userID, req.MenuID)
	if err != nil {
		handleOrderError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"order": order})
}

func (h OrderHandler) ListMine(c *gin.Context) {
	userID, ok := currentUserID(c)
	if !ok {
		return
	}

	orders, err := h.orders.ListByUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func (h OrderHandler) Cancel(c *gin.Context) {
	userID, ok := currentUserID(c)
	if !ok {
		return
	}
	orderID, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	order, err := h.orders.Cancel(c.Request.Context(), userID, orderID)
	if err != nil {
		handleOrderError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func (h OrderHandler) ListAll(c *gin.Context) {
	orders, err := h.orders.ListAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func currentUserID(c *gin.Context) (uint, bool) {
	userID, ok := c.Get(middleware.ContextUserID)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing auth context"})
		return 0, false
	}

	id, ok := userID.(uint)
	if !ok || id == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid auth context"})
		return 0, false
	}

	return id, true
}

func handleOrderError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, services.ErrDuplicateOrder):
		c.JSON(http.StatusConflict, gin.H{"error": "order already placed for this menu"})
	case errors.Is(err, services.ErrForbiddenOrder):
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
	case errors.Is(err, services.ErrOrderCancelled):
		c.JSON(http.StatusBadRequest, gin.H{"error": "order already cancelled"})
	case errors.Is(err, gorm.ErrRecordNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "order or menu not found"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "order operation failed"})
	}
}
