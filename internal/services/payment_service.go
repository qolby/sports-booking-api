package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/qolby/sports-booking-api/internal/models"
	"gorm.io/gorm"
)

type PaymentService struct {
	db *gorm.DB
}

func NewPaymentService(db *gorm.DB) *PaymentService {
	return &PaymentService{db: db}
}

type CreatePaymentRequest struct {
	BookingID     uint   `json:"booking_id" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
}

func (s *PaymentService) ProcessPayment(req CreatePaymentRequest) (*models.Payment, error) {
	// Get booking
	var booking models.Booking
	if err := s.db.First(&booking, req.BookingID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("booking not found")
		}
		return nil, err
	}

	// Check if booking is already paid
	if booking.Status == models.StatusPaid {
		return nil, errors.New("booking is already paid")
	}

	// Check if payment already exists
	var existingPayment models.Payment
	if err := s.db.Where("booking_id = ?", req.BookingID).First(&existingPayment).Error; err == nil {
		return nil, errors.New("payment already exists for this booking")
	}

	// Mock payment processing
	transactionID := fmt.Sprintf("TRX-%d-%d", booking.ID, time.Now().Unix())

	// Create payment
	payment := models.Payment{
		BookingID:     req.BookingID,
		Amount:        booking.TotalPrice,
		Status:        models.PaymentCompleted,
		PaymentMethod: req.PaymentMethod,
		TransactionID: transactionID,
	}

	// Start transaction
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create payment
	if err := tx.Create(&payment).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Update booking status
	if err := tx.Model(&booking).Update("status", models.StatusPaid).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Load relations
	s.db.Preload("Booking.Field").Preload("Booking.User").First(&payment, payment.ID)

	return &payment, nil
}
