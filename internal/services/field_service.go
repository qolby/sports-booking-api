package services

import (
	"errors"

	"github.com/qolby/sports-booking-api/internal/models"
	"gorm.io/gorm"
)

type FieldService struct {
	db *gorm.DB
}

func NewFieldService(db *gorm.DB) *FieldService {
	return &FieldService{db: db}
}

type CreateFieldRequest struct {
	Name         string `json:"name" validate:"required"`
	PricePerHour int    `json:"price_per_hour" validate:"required,gt=0"`
	Location     string `json:"location" validate:"required"`
}

type UpdateFieldRequest struct {
	Name         string `json:"name"`
	PricePerHour int    `json:"price_per_hour"`
	Location     string `json:"location"`
}

func (s *FieldService) CreateField(req CreateFieldRequest) (*models.Field, error) {
	field := models.Field{
		Name:         req.Name,
		PricePerHour: req.PricePerHour,
		Location:     req.Location,
	}

	if err := s.db.Create(&field).Error; err != nil {
		return nil, err
	}

	return &field, nil
}

func (s *FieldService) GetAllFields() ([]models.Field, error) {
	var fields []models.Field
	if err := s.db.Find(&fields).Error; err != nil {
		return nil, err
	}
	return fields, nil
}

func (s *FieldService) GetFieldByID(id uint) (*models.Field, error) {
	var field models.Field
	if err := s.db.First(&field, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("field not found")
		}
		return nil, err
	}
	return &field, nil
}

func (s *FieldService) UpdateField(id uint, req UpdateFieldRequest) (*models.Field, error) {
	field, err := s.GetFieldByID(id)
	if err != nil {
		return nil, err
	}

	updates := map[string]interface{}{}
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.PricePerHour > 0 {
		updates["price_per_hour"] = req.PricePerHour
	}
	if req.Location != "" {
		updates["location"] = req.Location
	}

	if err := s.db.Model(field).Updates(updates).Error; err != nil {
		return nil, err
	}

	return field, nil
}

func (s *FieldService) DeleteField(id uint) error {
	result := s.db.Delete(&models.Field{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("field not found")
	}
	return nil
}
