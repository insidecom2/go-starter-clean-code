package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	cfgpkg "github.com/example/go-starter/internal/config"
	"github.com/example/go-starter/internal/logger"
	httpHandler "github.com/example/go-starter/internal/http/handler"
	httpMiddleware "github.com/example/go-starter/internal/http/middleware"
	userHandler "github.com/example/go-starter/internal/user/handler"
	userRepo "github.com/example/go-starter/internal/user/repo"
	userUsecase "github.com/example/go-starter/internal/user/usecase"
)

func main() {
	cfg := cfgpkg.Load()
	logg, err := logger.New(cfg.LogLevel)
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	defer func() { _ = logg.Sync() }()

	app := fiber.New()

	// Middlewares
	app.Use(httpMiddleware.RequestID())
	app.Use(recover.New())

	// Domain wiring (simple DI)
	repo := userRepo.NewMemoryRepo()
	uc := userUsecase.New(repo)
	uh := userHandler.New(uc)

	api := app.Group("/api/v1")
	api.Get("/health", httpHandler.Health)

	users := api.Group("/users")
	users.Get("/", uh.List)
	users.Post("/", uh.Create)

	logg.Sugar().Infof("starting server on :%s", cfg.AppPort)
	if err := app.Listen(":" + cfg.AppPort); err != nil {
		logg.Sugar().Fatalf("server failed: %v", err)
	}
}
