package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	routerpkg "github.com/yourname/go-starter/scaffold/internal/adapter/http"
	"github.com/yourname/go-starter/scaffold/internal/config"
	"github.com/yourname/go-starter/scaffold/internal/middleware"
)

// New builds the fiber app with middleware and routes applied
func New(cfg config.Config, logger *zap.Logger, db *sqlx.DB) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ReadTimeout:           5 * time.Second,
		WriteTimeout:          10 * time.Second,
		IdleTimeout:           30 * time.Second,
	})

	app.Use(middleware.RequestID())
	app.Use(middleware.ZapLogger(logger))

	routerpkg.Setup(app, logger, db)
	return app
}
