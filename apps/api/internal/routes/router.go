package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/officebite/officebite/apps/api/internal/config"
	"github.com/officebite/officebite/apps/api/internal/handlers"
)

func NewRouter(cfg config.Config) *gin.Engine {
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.WebOrigin},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	healthHandler := handlers.NewHealthHandler()
	router.GET("/healthz", healthHandler.Check)

	api := router.Group("/api/v1")
	api.GET("/health", healthHandler.Check)

	return router
}
