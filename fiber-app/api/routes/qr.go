package routes

import (
	"fiber-app/api/handlers"
	"fiber-app/internal/domain/users"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRouter(app *fiber.App, userService users.UserService) {
	handler := handlers.NewUserHandler(userService)
	app.Post("/fiber/login", handler.Login)
}

func SetupHealthCheckRouter(app *fiber.App) {
	app.Get("/fiber/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
