package domain_test

import (
	"fiber-app/internal/domain/qr"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAdapter struct {
	mock.Mock
}

func (m *MockAdapter) QRFactorize(matrix [][]float64) ([][]float64, [][]float64, error) {
	args := m.Called(matrix)
	return args.Get(0).([][]float64), args.Get(1).([][]float64), args.Error(2)
}

func TestQRService(t *testing.T) {
	mockAdapter := new(MockAdapter)
	service := qr.NewService(mockAdapter)

	t.Run("Delega correctamente al adaptador", func(t *testing.T) {
		testMatrix := [][]float64{{1, 2}, {3, 4}}
		expectedQ := [][]float64{{0.1, 0.2}, {0.3, 0.4}}
		expectedR := [][]float64{{1.1, 1.2}, {1.3, 1.4}}

		mockAdapter.On("QRFactorize", testMatrix).Return(expectedQ, expectedR, nil)

		q, r, err := service.Factorize(testMatrix)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, expectedQ, q)
		assert.Equal(t, expectedR, r)
		mockAdapter.AssertExpectations(t)
	})
}
