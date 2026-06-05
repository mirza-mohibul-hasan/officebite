package services

import (
	"context"
	"errors"

	"github.com/officebite/officebite/apps/api/internal/models"
	"github.com/officebite/officebite/apps/api/internal/repository"
	"gorm.io/gorm"
)

var (
	ErrDuplicateOrder = errors.New("active order already exists")
	ErrForbiddenOrder = errors.New("order does not belong to user")
	ErrOrderCancelled = errors.New("order already cancelled")
)

type OrderService struct {
	orders repository.OrderRepository
	menus  repository.MenuRepository
}

func NewOrderService(orders repository.OrderRepository, menus repository.MenuRepository) OrderService {
	return OrderService{orders: orders, menus: menus}
}

func (s OrderService) Place(ctx context.Context, userID uint, menuID uint) (*models.Order, error) {
	if _, err := s.menus.FindByID(ctx, menuID); err != nil {
		return nil, err
	}

	existing, err := s.orders.FindActiveByUserAndMenu(ctx, userID, menuID)
	if err == nil && existing != nil {
		return nil, ErrDuplicateOrder
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	order := &models.Order{
		UserID: userID,
		MenuID: menuID,
		Status: models.OrderStatusPlaced,
	}
	if err := s.orders.Create(ctx, order); err != nil {
		return nil, err
	}

	return s.orders.FindByID(ctx, order.ID)
}

func (s OrderService) Cancel(ctx context.Context, userID uint, orderID uint) (*models.Order, error) {
	order, err := s.orders.FindByID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	if order.UserID != userID {
		return nil, ErrForbiddenOrder
	}
	if order.Status == models.OrderStatusCancelled {
		return nil, ErrOrderCancelled
	}

	order.Status = models.OrderStatusCancelled
	if err := s.orders.Update(ctx, order); err != nil {
		return nil, err
	}

	return s.orders.FindByID(ctx, order.ID)
}

func (s OrderService) ListByUser(ctx context.Context, userID uint) ([]models.Order, error) {
	return s.orders.ListByUser(ctx, userID)
}

func (s OrderService) ListAll(ctx context.Context) ([]models.Order, error) {
	return s.orders.ListAll(ctx)
}
