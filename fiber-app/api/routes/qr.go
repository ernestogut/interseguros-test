package routes

import (
	"fiber-app/api/handlers"
	"fiber-app/internal/domain/qr"
	"fiber-app/internal/domain/users"
	"fiber-app/internal/infrastructure/httpclient"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(qrService *qr.Service, client *httpclient.Client) *fiber.App {
	app := fiber.New()
	handler := handlers.NewQRHandler(qrService, client)

	app.Post("/process", handler.ProcessQR)
	return app
}

func SetupUserRouter(app *fiber.App, userService users.UserService) {
	handler := handlers.NewUserHandler(userService)
	app.Post("/login", handler.Login)
}
