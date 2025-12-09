package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/qolby/sports-booking-api/internal/services"
	"github.com/qolby/sports-booking-api/internal/utils"
)

type BookingHandler struct {
	bookingService *services.BookingService
}

func NewBookingHandler(bookingService *services.BookingService) *BookingHandler {
	return &BookingHandler{bookingService: bookingService}
}

func (h *BookingHandler) CreateBooking(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var req services.CreateBookingRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	booking, err := h.bookingService.CreateBooking(userID, req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to create booking", err)
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Booking created successfully", booking)
}

func (h *BookingHandler) GetUserBookings(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	bookings, err := h.bookingService.GetUserBookings(userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch bookings", err)
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Bookings retrieved successfully", bookings)
}

func (h *BookingHandler) GetBookingByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid booking ID", err)
	}

	booking, err := h.bookingService.GetBookingByID(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Booking not found", err)
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Booking retrieved successfully", booking)
}
