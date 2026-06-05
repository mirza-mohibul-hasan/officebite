package repository

import (
	"context"
	"time"

	"github.com/officebite/officebite/apps/api/internal/models"
	"gorm.io/gorm"
)

type DashboardSummary struct {
	TodayOrders     int64 `json:"today_orders"`
	PlacedOrders    int64 `json:"placed_orders"`
	CancelledOrders int64 `json:"cancelled_orders"`
	TodayMenus      int64 `json:"today_menus"`
	TodayRevenue    int64 `json:"today_revenue"`
}

type AnalyticsRepository interface {
	Summary(ctx context.Context, today time.Time) (DashboardSummary, error)
}

type GormAnalyticsRepository struct {
	db *gorm.DB
}

func NewAnalyticsRepository(db *gorm.DB) AnalyticsRepository {
	return &GormAnalyticsRepository{db: db}
}

func (r *GormAnalyticsRepository) Summary(ctx context.Context, today time.Time) (DashboardSummary, error) {
	var summary DashboardSummary
	start := beginningOfDay(today)
	end := start.AddDate(0, 0, 1)

	if err := r.db.WithContext(ctx).Model(&models.Order{}).
		Where("created_at >= ? AND created_at < ?", start, end).
		Count(&summary.TodayOrders).Error; err != nil {
		return summary, err
	}
	if err := r.db.WithContext(ctx).Model(&models.Order{}).
		Where("status = ?", models.OrderStatusPlaced).
		Count(&summary.PlacedOrders).Error; err != nil {
		return summary, err
	}
	if err := r.db.WithContext(ctx).Model(&models.Order{}).
		Where("status = ?", models.OrderStatusCancelled).
		Count(&summary.CancelledOrders).Error; err != nil {
		return summary, err
	}
	if err := r.db.WithContext(ctx).Model(&models.Menu{}).
		Where("available_date = ?", today.Format("2006-01-02")).
		Count(&summary.TodayMenus).Error; err != nil {
		return summary, err
	}
	if err := r.db.WithContext(ctx).Model(&models.Order{}).
		Select("COALESCE(SUM(menus.price), 0)").
		Joins("JOIN menus ON menus.id = orders.menu_id").
		Where("orders.status = ? AND orders.created_at >= ? AND orders.created_at < ?", models.OrderStatusPlaced, start, end).
		Scan(&summary.TodayRevenue).Error; err != nil {
		return summary, err
	}

	return summary, nil
}

func beginningOfDay(value time.Time) time.Time {
	year, month, day := value.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, value.Location())
}
