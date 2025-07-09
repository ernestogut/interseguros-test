package qr

import (
	"errors"
	"fiber-app/internal/infrastructure/math"
)

type QRService interface {
	Factorize(data [][]float64) (q, r [][]float64, err error)
}

type Service struct {
	matrixAdapter math.MatrixAdapter
}

func NewService(adapter math.MatrixAdapter) *Service {
	return &Service{matrixAdapter: adapter}
}

func (s *Service) Factorize(data [][]float64) (q, r [][]float64, err error) {
	if len(data) == 0 || len(data[0]) == 0 {
		return nil, nil, ErrEmptyMatrix
	}

	// Delegar a la infraestructura matemática
	return s.matrixAdapter.QRFactorize(data)
}

var ErrEmptyMatrix = errors.New("matrix cannot be empty")
