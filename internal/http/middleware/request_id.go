package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// RequestID adds a unique request id header and stores it in locals
func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := uuid.New().String()
		c.Set("X-Request-ID", id)
		c.Locals("requestid", id)
		return c.Next()
	}
}
