package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Options struct {
	DatabaseURL string
	LogLevel    logger.LogLevel
}

func Connect(options Options) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(options.DatabaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(options.LogLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("open database connection: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("read database handle: %w", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return db, nil
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("read database handle: %w", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("close database connection: %w", err)
	}

	return nil
}
