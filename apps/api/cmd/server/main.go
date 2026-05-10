package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/officebite/officebite/apps/api/internal/config"
	"github.com/officebite/officebite/apps/api/internal/database"
	"github.com/officebite/officebite/apps/api/internal/repository"
	"github.com/officebite/officebite/apps/api/internal/routes"
	"gorm.io/gorm/logger"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(database.Options{
		DatabaseURL: cfg.DatabaseURL,
		LogLevel:    databaseLogLevel(cfg.DatabaseLogLevel),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer func() {
		if err := database.Close(db); err != nil {
			log.Printf("failed to close database: %v", err)
		}
	}()

	if cfg.AutoMigrate {
		if err := database.AutoMigrate(db); err != nil {
			log.Fatalf("failed to run database migrations: %v", err)
		}
	}

	repositories := repository.NewRepositories(db)
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to read database handle: %v", err)
	}

	router := routes.NewRouter(cfg, routes.Dependencies{
		SQLDB:        sqlDB,
		Repositories: repositories,
	})

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("officebite api listening on :%s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start api server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown api server: %v", err)
	}
}

func databaseLogLevel(value string) logger.LogLevel {
	switch value {
	case "silent":
		return logger.Silent
	case "info":
		return logger.Info
	case "error":
		return logger.Error
	default:
		return logger.Warn
	}
}
