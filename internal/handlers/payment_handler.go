package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qolby/sports-booking-api/internal/services"
	"github.com/qolby/sports-booking-api/internal/utils"
)

type PaymentHandler struct {
	paymentService *services.PaymentService
}

func NewPaymentHandler(paymentService *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (h *PaymentHandler) ProcessPayment(c *fiber.Ctx) error {
	var req services.CreatePaymentRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	payment, err := h.paymentService.ProcessPayment(req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Payment processing failed", err)
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Payment processed successfully", payment)
}
