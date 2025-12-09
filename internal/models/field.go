package models

import (
	"time"

	"gorm.io/gorm"
)

type Field struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	Name         string         `gorm:"not null" json:"name"`
	PricePerHour int            `gorm:"not null" json:"price_per_hour"`
	Location     string         `gorm:"not null" json:"location"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Bookings     []Booking      `gorm:"foreignKey:FieldID" json:"bookings,omitempty"`
}
