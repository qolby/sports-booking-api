package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/qolby/sports-booking-api/internal/services"
	"github.com/qolby/sports-booking-api/internal/utils"
)

type FieldHandler struct {
	fieldService *services.FieldService
}

func NewFieldHandler(fieldService *services.FieldService) *FieldHandler {
	return &FieldHandler{fieldService: fieldService}
}

func (h *FieldHandler) CreateField(c *fiber.Ctx) error {
	var req services.CreateFieldRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	field, err := h.fieldService.CreateField(req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create field", err)
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Field created successfully", field)
}

func (h *FieldHandler) GetAllFields(c *fiber.Ctx) error {
	fields, err := h.fieldService.GetAllFields()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch fields", err)
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Fields retrieved successfully", fields)
}

func (h *FieldHandler) GetFieldByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid field ID", err)
	}

	field, err := h.fieldService.GetFieldByID(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Field not found", err)
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Field retrieved successfully", field)
}

func (h *FieldHandler) UpdateField(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid field ID", err)
	}

	var req services.UpdateFieldRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	field, err := h.fieldService.UpdateField(uint(id), req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to update field", err)
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Field updated successfully", field)
}

func (h *FieldHandler) DeleteField(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid field ID", err)
	}

	if err := h.fieldService.DeleteField(uint(id)); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete field", err)
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Field deleted successfully", nil)
}
