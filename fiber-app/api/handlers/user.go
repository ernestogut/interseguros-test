package handlers

import (
	"fiber-app/internal/domain/users"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service users.UserService
}

func NewUserHandler(s users.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req users.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request format"})
	}
	token, err := h.service.Login(req)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	return c.JSON(fiber.Map{"token": token})
}
