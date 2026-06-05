package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/officebite/officebite/apps/api/internal/services"
)

type AnalyticsHandler struct {
	analytics services.AnalyticsService
}

func NewAnalyticsHandler(analytics services.AnalyticsService) AnalyticsHandler {
	return AnalyticsHandler{analytics: analytics}
}

func (h AnalyticsHandler) DashboardSummary(c *gin.Context) {
	summary, err := h.analytics.DashboardSummary(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load dashboard summary"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"summary": summary})
}
