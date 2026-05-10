package database

import (
	"fmt"

	"github.com/officebite/officebite/apps/api/internal/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.User{},
		&models.Menu{},
		&models.Order{},
	); err != nil {
		return fmt.Errorf("auto migrate database: %w", err)
	}

	return nil
}
