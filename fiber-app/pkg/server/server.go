package server

import (
	"fiber-app/api/handlers"
	"fiber-app/api/middleware"
	"fiber-app/api/routes"
	"fiber-app/internal/domain/qr"
	"fiber-app/internal/domain/users"
	"fiber-app/internal/infrastructure/httpclient"
	"fiber-app/internal/infrastructure/math"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

func NewApp() *fiber.App {
	godotenv.Load()
	matrixAdapter := math.NewRealMatrixAdapter()
	qrService := qr.NewService(matrixAdapter)
	nodeJSClient := httpclient.New(os.Getenv("NODE_APP_URL"))
	userService := users.NewService()

	app := fiber.New()
	// Solo login esta libre
	routes.SetupHealthCheckRouter(app)
	routes.SetupUserRouter(app, userService)
	// Rutas protegidas
	protected := app.Group("", middleware.ValidateJWTMiddleware)
	qrHandler := handlers.NewQRHandler(qrService, nodeJSClient)
	protected.Post("/fiber/process", qrHandler.ProcessQR)
	return app
}

func NewTestApp() *fiber.App {
	return NewApp()
}
