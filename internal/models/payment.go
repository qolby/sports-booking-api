package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "pending"
	PaymentCompleted PaymentStatus = "paid"
	PaymentFailed    PaymentStatus = "failed"
)

type Payment struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	BookingID     uint           `gorm:"uniqueIndex;not null" json:"booking_id"`
	Booking       Booking        `gorm:"foreignKey:BookingID" json:"booking,omitempty"`
	Amount        int            `gorm:"not null" json:"amount"`
	Status        PaymentStatus  `gorm:"type:varchar(20);default:'pending'" json:"status"`
	PaymentMethod string         `json:"payment_method"`
	TransactionID string         `gorm:"uniqueIndex" json:"transaction_id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
