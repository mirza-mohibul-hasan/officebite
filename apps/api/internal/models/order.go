package models

import "time"

type OrderStatus string

const (
	OrderStatusPlaced    OrderStatus = "placed"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    uint        `gorm:"not null;index" json:"user_id"`
	MenuID    uint        `gorm:"not null;index" json:"menu_id"`
	Status    OrderStatus `gorm:"type:varchar(24);not null;default:'placed';index" json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`

	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"user,omitempty"`
	Menu Menu `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"menu,omitempty"`
}
