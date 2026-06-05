package routes

import (
	"database/sql"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/officebite/officebite/apps/api/internal/config"
	"github.com/officebite/officebite/apps/api/internal/handlers"
	"github.com/officebite/officebite/apps/api/internal/middleware"
	"github.com/officebite/officebite/apps/api/internal/repository"
	"github.com/officebite/officebite/apps/api/internal/services"
)

type Dependencies struct {
	SQLDB        *sql.DB
	Repositories repository.Repositories
}

func NewRouter(cfg config.Config, deps Dependencies) *gin.Engine {
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

	healthHandler := handlers.NewHealthHandler(deps.SQLDB)
	router.GET("/healthz", healthHandler.Check)
	router.GET("/readyz", healthHandler.Ready)

	api := router.Group("/api/v1")
	api.GET("/health", healthHandler.Check)
	api.GET("/ready", healthHandler.Ready)

	authService := services.NewAuthService(deps.Repositories.Users, cfg.JWTSecret, cfg.JWTIssuer)
	authHandler := handlers.NewAuthHandler(authService)
	api.POST("/auth/login", authHandler.Login)

	protected := api.Group("")
	protected.Use(middleware.AuthRequired(cfg.JWTSecret))
	protected.GET("/auth/me", authHandler.Me)

	menuService := services.NewMenuService(deps.Repositories.Menus)
	menuHandler := handlers.NewMenuHandler(menuService)
	protected.GET("/menus/today", menuHandler.ListToday)

	admin := protected.Group("/admin")
	admin.Use(middleware.RequireRole("admin"))
	admin.GET("/menus", menuHandler.ListAll)
	admin.POST("/menus", menuHandler.Create)
	admin.PUT("/menus/:id", menuHandler.Update)
	admin.DELETE("/menus/:id", menuHandler.Delete)

	return router
}
