package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qolby/sports-booking-api/internal/services"
	"github.com/qolby/sports-booking-api/internal/utils"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user account with email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body services.RegisterRequest true "Registration details"
// @Success 201 {object} utils.Response{data=services.AuthResponse}
// @Failure 400 {object} utils.Response
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req services.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	result, err := h.authService.Register(req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Registration failed", err)
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Registration successful", result)
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body services.LoginRequest true "Login credentials"
// @Success 200 {object} utils.Response{data=services.AuthResponse}
// @Failure 401 {object} utils.Response
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req services.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body", err)
	}

	result, err := h.authService.Login(req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Login failed", err)
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Login successful", result)
}
