package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/example/go-starter/internal/handler"
)

// Setup registers routes and route groups
func Setup(app *fiber.App, logger *zap.SugaredLogger, db *sqlx.DB) {
	api := app.Group("/api")
	api.Get("/health", handler.Health)
	// add application routes here (users, auth, etc.)
}
