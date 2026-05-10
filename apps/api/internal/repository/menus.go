package repository

import (
	"context"
	"time"

	"github.com/officebite/officebite/apps/api/internal/models"
	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(ctx context.Context, menu *models.Menu) error
	Update(ctx context.Context, menu *models.Menu) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*models.Menu, error)
	ListByDate(ctx context.Context, date time.Time) ([]models.Menu, error)
}

type GormMenuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &GormMenuRepository{db: db}
}

func (r *GormMenuRepository) Create(ctx context.Context, menu *models.Menu) error {
	return r.db.WithContext(ctx).Create(menu).Error
}

func (r *GormMenuRepository) Update(ctx context.Context, menu *models.Menu) error {
	return r.db.WithContext(ctx).Save(menu).Error
}

func (r *GormMenuRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Menu{}, id).Error
}

func (r *GormMenuRepository) FindByID(ctx context.Context, id uint) (*models.Menu, error) {
	var menu models.Menu
	if err := r.db.WithContext(ctx).First(&menu, id).Error; err != nil {
		return nil, err
	}

	return &menu, nil
}

func (r *GormMenuRepository) ListByDate(ctx context.Context, date time.Time) ([]models.Menu, error) {
	var menus []models.Menu
	if err := r.db.WithContext(ctx).
		Where("available_date = ?", date.Format("2006-01-02")).
		Order("title ASC").
		Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}
