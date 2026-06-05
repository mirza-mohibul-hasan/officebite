package repository

import (
	"context"

	"github.com/officebite/officebite/apps/api/internal/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, order *models.Order) error
	Update(ctx context.Context, order *models.Order) error
	FindByID(ctx context.Context, id uint) (*models.Order, error)
	FindActiveByUserAndMenu(ctx context.Context, userID uint, menuID uint) (*models.Order, error)
	CountActiveByMenu(ctx context.Context, menuID uint) (int64, error)
	ListByUser(ctx context.Context, userID uint) ([]models.Order, error)
	ListAll(ctx context.Context) ([]models.Order, error)
}

type GormOrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) Create(ctx context.Context, order *models.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *GormOrderRepository) Update(ctx context.Context, order *models.Order) error {
	return r.db.WithContext(ctx).Save(order).Error
}

func (r *GormOrderRepository) FindByID(ctx context.Context, id uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.WithContext(ctx).Preload("User").Preload("Menu").First(&order, id).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *GormOrderRepository) FindActiveByUserAndMenu(ctx context.Context, userID uint, menuID uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.WithContext(ctx).
		Where("user_id = ? AND menu_id = ? AND status IN ?", userID, menuID, []models.OrderStatus{models.OrderStatusPlaced, models.OrderStatusConfirmed}).
		First(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *GormOrderRepository) CountActiveByMenu(ctx context.Context, menuID uint) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&models.Order{}).
		Where("menu_id = ? AND status IN ?", menuID, []models.OrderStatus{models.OrderStatusPlaced, models.OrderStatusConfirmed}).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (r *GormOrderRepository) ListByUser(ctx context.Context, userID uint) ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.WithContext(ctx).
		Preload("Menu").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *GormOrderRepository) ListAll(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Menu").
		Order("created_at DESC").
		Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}
