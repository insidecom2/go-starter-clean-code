package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// RequestID ensures each request has an id available via X-Request-Id header
func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Get("X-Request-Id")
		if id == "" {
			id = uuid.New().String()
			c.Set("X-Request-Id", id)
		}
		c.Locals("requestid", id)
		return c.Next()
	}
}

// ZapLogger logs requests using a zap logger
func ZapLogger(logger *zap.SugaredLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		latency := time.Since(start)
		reqID, _ := c.Locals("requestid").(string)
		logger.Infow("request",
			"method", c.Method(),
			"path", c.Path(),
			"request_id", reqID,
			"status", c.Response().StatusCode(),
			"latency", latency.String(),
		)
		return err
	}
}
