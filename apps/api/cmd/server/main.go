package main

import (
	"log"

	"github.com/officebite/officebite/apps/api/internal/config"
	"github.com/officebite/officebite/apps/api/internal/routes"
)

func main() {
	cfg := config.Load()
	router := routes.NewRouter(cfg)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("failed to start api server: %v", err)
	}
}
