package services

import (
	"testing"
	"time"

	"github.com/qolby/sports-booking-api/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupBookingTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}

	db.AutoMigrate(&models.User{}, &models.Field{}, &models.Booking{})

	return db
}

func TestBookingService_CreateBooking(t *testing.T) {
	db := setupBookingTestDB()
	bookingService := NewBookingService(db)

	// Create test data
	user := models.User{Email: "test@example.com", Name: "Test User", Role: models.RoleUser}
	user.HashPassword("password")
	db.Create(&user)

	field := models.Field{
		Name:         "Test Field",
		PricePerHour: 100000,
		Location:     "Test Location",
	}
	db.Create(&field)

	now := time.Now()
	startTime := now.Add(24 * time.Hour)
	endTime := startTime.Add(2 * time.Hour)

	tests := []struct {
		name    string
		request CreateBookingRequest
		wantErr bool
	}{
		{
			name: "Successful booking",
			request: CreateBookingRequest{
				FieldID:   field.ID,
				StartTime: startTime,
				EndTime:   endTime,
			},
			wantErr: false,
		},
		{
			name: "End time before start time",
			request: CreateBookingRequest{
				FieldID:   field.ID,
				StartTime: endTime,
				EndTime:   startTime,
			},
			wantErr: true,
		},
		{
			name: "Overlapping booking",
			request: CreateBookingRequest{
				FieldID:   field.ID,
				StartTime: startTime.Add(30 * time.Minute),
				EndTime:   endTime.Add(30 * time.Minute),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := bookingService.CreateBooking(user.ID, tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, user.ID, result.UserID)
				assert.Equal(t, field.ID, result.FieldID)
				assert.Equal(t, models.StatusPending, result.Status)
			}
		})
	}
}
