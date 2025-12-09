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

// CreateField godoc
// @Summary Create a new field
// @Description Create a new sports field (Admin only)
// @Tags Fields
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body services.CreateFieldRequest true "Field details"
// @Success 201 {object} utils.Response{data=models.Field}
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 403 {object} utils.Response
// @Router /fields [post]
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

// GetAllFields godoc
// @Summary Get all fields
// @Description Get list of all available sports fields
// @Tags Fields
// @Produce json
// @Success 200 {object} utils.Response{data=[]models.Field}
// @Router /fields [get]
func (h *FieldHandler) GetAllFields(c *fiber.Ctx) error {
	fields, err := h.fieldService.GetAllFields()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch fields", err)
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Fields retrieved successfully", fields)
}

// GetFieldByID godoc
// @Summary Get field by ID
// @Description Get details of a specific field
// @Tags Fields
// @Produce json
// @Param id path int true "Field ID"
// @Success 200 {object} utils.Response{data=models.Field}
// @Failure 404 {object} utils.Response
// @Router /fields/{id} [get]
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

// UpdateField godoc
// @Summary Update field
// @Description Update field details (Admin only)
// @Tags Fields
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Field ID"
// @Param request body services.UpdateFieldRequest true "Update details"
// @Success 200 {object} utils.Response{data=models.Field}
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 403 {object} utils.Response
// @Router /fields/{id} [put]
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

// DeleteField godoc
// @Summary Delete field
// @Description Delete a field (Admin only)
// @Tags Fields
// @Security BearerAuth
// @Param id path int true "Field ID"
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 403 {object} utils.Response
// @Router /fields/{id} [delete]
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
