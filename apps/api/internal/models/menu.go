package models

import (
	"time"
)

type Menu struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Title         string    `gorm:"size:160;not null" json:"title"`
	Description   string    `gorm:"type:text;not null" json:"description"`
	Price         int64     `gorm:"not null" json:"price"`
	AvailableDate time.Time `gorm:"type:date;not null;index" json:"available_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
