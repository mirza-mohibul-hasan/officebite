package models

import (
	"time"
)

type Menu struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Title         string    `gorm:"size:160;not null" json:"title"`
	Description   string    `gorm:"type:text;not null" json:"description"`
	Category      string    `gorm:"size:80;not null;default:'lunch'" json:"category"`
	Price         int64     `gorm:"not null" json:"price"`
	AvailableDate time.Time `gorm:"type:date;not null;index" json:"available_date"`
	CutoffTime    time.Time `gorm:"not null" json:"cutoff_time"`
	MaxOrders     int       `gorm:"not null;default:0" json:"max_orders"`
	IsActive      bool      `gorm:"not null;default:true" json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
