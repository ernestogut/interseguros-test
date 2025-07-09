package users

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Login(request LoginRequest) (string, error) {
	if request.Username != "admin" || request.Password != "admin" {
		return "", errors.New("invalid credentials")
	}
	claims := jwt.MapClaims{
		"username": request.Username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}

type UserService interface {
	Login(request LoginRequest) (string, error)
}

var _ UserService = (*Service)(nil)
