package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yourname/go-starter/scaffold/internal/user/usecase"
)

// Handlers depend on usecases; wiring happens in router

func CreateUser(u *usecase.Usecase) fiber.Handler {
    type req struct{ Email, Password string }
    return func(c *fiber.Ctx) error {
        var r req
        if err := c.BodyParser(&r); err != nil {
            return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
        }
        user, err := u.CreateUser(c.UserContext(), r.Email, r.Password)
        if err != nil {
            return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        }
        return c.Status(http.StatusCreated).JSON(fiber.Map{"id": user.ID})
    }
}

func GetUser(u *usecase.Usecase) fiber.Handler {
    return func(c *fiber.Ctx) error {
        idStr := c.Params("id")
        id, err := uuid.Parse(idStr)
        if err != nil { return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error":"invalid id"}) }
        user, err := u.GetUser(c.UserContext(), id)
        if err != nil { return c.Status(http.StatusNotFound).JSON(fiber.Map{"error":"not found"}) }
        return c.JSON(fiber.Map{"id": user.ID, "email": user.Email, "created_at": user.CreatedAt})
    }
}

func Authenticate(u *usecase.Usecase) fiber.Handler {
    type req struct{ Email, Password string }
    return func(c *fiber.Ctx) error {
        var r req
        if err := c.BodyParser(&r); err != nil { return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"}) }
        user, err := u.Authenticate(c.UserContext(), r.Email, r.Password)
        if err != nil { return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error":"invalid creds"}) }
        // For demo return user id; real app should return JWT
        return c.JSON(fiber.Map{"id": user.ID})
    }
}
