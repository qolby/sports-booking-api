package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qolby/sports-booking-api/internal/utils"
)

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("userRole").(string)

		if role != "admin" {
			return utils.ErrorResponse(c, fiber.StatusForbidden, "Admin access required", nil)
		}

		return c.Next()
	}
}
