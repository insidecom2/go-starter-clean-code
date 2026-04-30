package handler

import (
	"github.com/google/uuid"
	"github.com/gofiber/fiber/v2"

	"github.com/example/go-starter/internal/user"
	"github.com/example/go-starter/internal/user/usecase"
)

// Handler exposes HTTP endpoints for users
type Handler struct {
	uc *usecase.UserUsecase
}

func New(uc *usecase.UserUsecase) *Handler { return &Handler{uc: uc} }

func (h *Handler) List(c *fiber.Ctx) error {
	return c.JSON(h.uc.List())
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var u user.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	h.uc.Create(u)
	return c.Status(fiber.StatusCreated).JSON(u)
}
