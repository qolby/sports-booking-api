package services

import (
	"errors"
	"time"

	"github.com/qolby/sports-booking-api/internal/models"
	"gorm.io/gorm"
)

type BookingService struct {
	db *gorm.DB
}

func NewBookingService(db *gorm.DB) *BookingService {
	return &BookingService{db: db}
}

type CreateBookingRequest struct {
	FieldID   uint      `json:"field_id" validate:"required"`
	StartTime time.Time `json:"start_time" validate:"required"`
	EndTime   time.Time `json:"end_time" validate:"required"`
}

func (s *BookingService) CreateBooking(userID uint, req CreateBookingRequest) (*models.Booking, error) {
	// Validate time range
	if req.EndTime.Before(req.StartTime) || req.EndTime.Equal(req.StartTime) {
		return nil, errors.New("end time must be after start time")
	}

	// Check if field exists
	var field models.Field
	if err := s.db.First(&field, req.FieldID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("field not found")
		}
		return nil, err
	}

	// Check for overlapping bookings
	var count int64
	s.db.Model(&models.Booking{}).Where(
		"field_id = ? AND status != ? AND ((start_time < ? AND end_time > ?) OR (start_time < ? AND end_time > ?) OR (start_time >= ? AND end_time <= ?))",
		req.FieldID,
		models.StatusCancelled,
		req.EndTime, req.StartTime,
		req.EndTime, req.EndTime,
		req.StartTime, req.EndTime,
	).Count(&count)

	if count > 0 {
		return nil, errors.New("field is already booked for this time slot")
	}

	// Calculate total price
	duration := req.EndTime.Sub(req.StartTime).Hours()
	totalPrice := int(duration * float64(field.PricePerHour))

	booking := models.Booking{
		UserID:     userID,
		FieldID:    req.FieldID,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		Status:     models.StatusPending,
		TotalPrice: totalPrice,
	}

	if err := s.db.Create(&booking).Error; err != nil {
		return nil, err
	}

	// Load relations
	s.db.Preload("Field").Preload("User").First(&booking, booking.ID)

	return &booking, nil
}

func (s *BookingService) GetUserBookings(userID uint) ([]models.Booking, error) {
	var bookings []models.Booking
	if err := s.db.Preload("Field").Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (s *BookingService) GetBookingByID(id uint) (*models.Booking, error) {
	var booking models.Booking
	if err := s.db.Preload("Field").Preload("User").Preload("Payment").First(&booking, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("booking not found")
		}
		return nil, err
	}
	return &booking, nil
}
