package services

import (
	"context"
	"time"

	"github.com/officebite/officebite/apps/api/internal/repository"
)

type AnalyticsService struct {
	analytics repository.AnalyticsRepository
}

func NewAnalyticsService(analytics repository.AnalyticsRepository) AnalyticsService {
	return AnalyticsService{analytics: analytics}
}

func (s AnalyticsService) DashboardSummary(ctx context.Context) (repository.DashboardSummary, error) {
	return s.analytics.Summary(ctx, time.Now())
}
