package http

import (
    "github.com/gofiber/fiber/v2"
    "go.uber.org/zap"
    "github.com/jmoiron/sqlx"

    "github.com/yourname/go-starter/scaffold/internal/handler"
    userrepo "github.com/yourname/go-starter/scaffold/internal/user/repo"
    userusecase "github.com/yourname/go-starter/scaffold/internal/user/usecase"
)

// Setup registers routes and route groups
func Setup(app *fiber.App, logger *zap.Logger, db *sqlx.DB) {
    api := app.Group("/api")
    api.Get("/health", handler.Health)

    // user endpoints
    ur := userrepo.NewRepository(db)
    uu := userusecase.NewUsecase(ur)
    users := api.Group("/users")
    users.Post("/", handler.CreateUser(uu))
    users.Get(":id", handler.GetUser(uu))
    users.Post("/login", handler.Authenticate(uu))
}
