package handler

import "github.com/gofiber/fiber/v2"

// Health returns a minimal healthy response
func Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
}
