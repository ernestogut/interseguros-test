package handlers_test

import (
	"fiber-app/api/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type QRService interface {
	Factorize(matrix [][]float64) ([][]float64, [][]float64, error)
}

type MockService struct {
	mock.Mock
}

func (m *MockService) Factorize(matrix [][]float64) ([][]float64, [][]float64, error) {
	args := m.Called(matrix)
	return args.Get(0).([][]float64), args.Get(1).([][]float64), args.Error(2)
}

func TestQRHandler(t *testing.T) {
	mockService := new(MockService)
	handler := handlers.NewQRHandler(mockService, nil)

	app := fiber.New()
	app.Post("/fiber/process", handler.ProcessQR)

	t.Run("Valid request", func(t *testing.T) {
		mockService.On("Factorize", mock.Anything).Return(
			[][]float64{{1, 0}, {0, 1}},
			[][]float64{{2, 0}, {0, 3}},
			nil,
		)

		req := httptest.NewRequest(
			"POST",
			"/fiber/process",
			strings.NewReader(`{"data":[[1,2],[3,4]]}`),
		)
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
