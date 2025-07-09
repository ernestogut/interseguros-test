package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWTMiddleware(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if auth == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "No token provided")
	}

	tokenStr := auth[len("Bearer "):]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Locals("user", claims)

	return c.Next()
}
