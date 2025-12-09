package models

import (
	"time"

	"gorm.io/gorm"
)

type BookingStatus string

const (
	StatusPending   BookingStatus = "pending"
	StatusPaid      BookingStatus = "paid"
	StatusCancelled BookingStatus = "cancelled"
)

type Booking struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	UserID     uint           `gorm:"not null" json:"user_id"`
	User       User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	FieldID    uint           `gorm:"not null" json:"field_id"`
	Field      Field          `gorm:"foreignKey:FieldID" json:"field,omitempty"`
	StartTime  time.Time      `gorm:"not null" json:"start_time"`
	EndTime    time.Time      `gorm:"not null" json:"end_time"`
	Status     BookingStatus  `gorm:"type:varchar(20);default:'pending'" json:"status"`
	TotalPrice int            `json:"total_price"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Payment    *Payment       `gorm:"foreignKey:BookingID" json:"payment,omitempty"`
}
