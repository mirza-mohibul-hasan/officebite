package database

import (
	"fmt"

	"github.com/officebite/officebite/apps/api/internal/models"
	"github.com/officebite/officebite/apps/api/internal/utils"
	"gorm.io/gorm"
)

type SeedOptions struct {
	AdminEmail       string
	AdminPassword    string
	EmployeeEmail    string
	EmployeePassword string
}

func Seed(db *gorm.DB, options SeedOptions) error {
	users := []struct {
		name     string
		email    string
		password string
		role     models.UserRole
	}{
		{name: "Admin User", email: options.AdminEmail, password: options.AdminPassword, role: models.RoleAdmin},
		{name: "Employee User", email: options.EmployeeEmail, password: options.EmployeePassword, role: models.RoleEmployee},
	}

	for _, seedUser := range users {
		var count int64
		if err := db.Model(&models.User{}).Where("email = ?", seedUser.email).Count(&count).Error; err != nil {
			return fmt.Errorf("check seed user: %w", err)
		}
		if count > 0 {
			continue
		}

		hash, err := utils.HashPassword(seedUser.password)
		if err != nil {
			return fmt.Errorf("hash seed password: %w", err)
		}

		user := models.User{
			Name:         seedUser.name,
			Email:        seedUser.email,
			PasswordHash: hash,
			Role:         seedUser.role,
		}
		if err := db.Create(&user).Error; err != nil {
			return fmt.Errorf("create seed user: %w", err)
		}
	}

	return nil
}
