package services

import (
	"testing"

	"github.com/qolby/sports-booking-api/internal/config"
	"github.com/qolby/sports-booking-api/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}

	// Auto migrate
	db.AutoMigrate(&models.User{})

	return db
}

func TestAuthService_Register(t *testing.T) {
	db := setupTestDB()
	cfg := &config.Config{
		JWT: config.JWTConfig{
			Secret: "test-secret",
			Expiry: 24,
		},
	}
	authService := NewAuthService(db, cfg)

	tests := []struct {
		name    string
		request RegisterRequest
		wantErr bool
	}{
		{
			name: "Successful registration",
			request: RegisterRequest{
				Email:    "test@example.com",
				Password: "password123",
				Name:     "Test User",
			},
			wantErr: false,
		},
		{
			name: "Duplicate email",
			request: RegisterRequest{
				Email:    "test@example.com",
				Password: "password123",
				Name:     "Test User 2",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := authService.Register(tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.request.Email, result.User.Email)
				assert.NotEmpty(t, result.Token)
			}
		})
	}
}

func TestAuthService_Login(t *testing.T) {
	db := setupTestDB()
	cfg := &config.Config{
		JWT: config.JWTConfig{
			Secret: "test-secret",
			Expiry: 24,
		},
	}
	authService := NewAuthService(db, cfg)

	// First, register a user
	registerReq := RegisterRequest{
		Email:    "login@example.com",
		Password: "password123",
		Name:     "Login Test",
	}
	authService.Register(registerReq)

	tests := []struct {
		name    string
		request LoginRequest
		wantErr bool
	}{
		{
			name: "Successful login",
			request: LoginRequest{
				Email:    "login@example.com",
				Password: "password123",
			},
			wantErr: false,
		},
		{
			name: "Wrong password",
			request: LoginRequest{
				Email:    "login@example.com",
				Password: "wrongpassword",
			},
			wantErr: true,
		},
		{
			name: "User not found",
			request: LoginRequest{
				Email:    "notfound@example.com",
				Password: "password123",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := authService.Login(tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.request.Email, result.User.Email)
				assert.NotEmpty(t, result.Token)
			}
		})
	}
}
