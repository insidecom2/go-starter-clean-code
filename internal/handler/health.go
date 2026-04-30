package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Health returns a small payload used by health checks
func Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"time":   time.Now().UTC().Format(time.RFC3339),
	})
}
