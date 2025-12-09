package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/qolby/sports-booking-api/internal/config"
	"github.com/qolby/sports-booking-api/internal/utils"
)

func AuthRequired(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Missing authorization header", nil)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid authorization format", nil)
		}

		claims, err := utils.ValidateToken(tokenString, cfg)
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid or expired token", err)
		}

		// Store user info in context
		c.Locals("userID", claims.UserID)
		c.Locals("userEmail", claims.Email)
		c.Locals("userRole", claims.Role)

		return c.Next()
	}
}
