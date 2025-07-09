package math

import (
	"errors"
	"fiber-app/pkg/utils"

	"gonum.org/v1/gonum/mat"
)

type MatrixAdapter interface {
	QRFactorize(matrix [][]float64) (q, r [][]float64, err error)
}

// Implementación concreta (tu struct original)
type RealMatrixAdapter struct{}

func NewRealMatrixAdapter() *RealMatrixAdapter {
	return &RealMatrixAdapter{}
}

func (m *RealMatrixAdapter) QRFactorize(matrix [][]float64) (q, r [][]float64, err error) {
	rows := len(matrix)
	if rows == 0 {
		return nil, nil, ErrEmptyMatrix
	}
	cols := len(matrix[0])

	flat := m.flattenMatrix(matrix)
	A := mat.NewDense(rows, cols, flat)

	var qr mat.QR
	qr.Factorize(A)

	var Q, R mat.Dense
	qr.QTo(&Q)
	qr.RTo(&R)

	qData := m.denseToSlices(&Q)
	rData := m.denseToSlices(&R)

	m.ensurePositiveDiagonal(qData, rData)
	m.roundMatrix(qData, 6)
	m.roundMatrix(rData, 6)

	return qData, rData, nil
}

// flattenMatrix convierte matriz 2D a slice 1D para gonum
func (m *RealMatrixAdapter) flattenMatrix(matrix [][]float64) []float64 {
	var flat []float64
	for _, row := range matrix {
		flat = append(flat, row...)
	}
	return flat
}

// denseToSlices convierte mat.Dense a [][]float64
func (m *RealMatrixAdapter) denseToSlices(dense *mat.Dense) [][]float64 {
	rows, cols := dense.Dims()
	slices := make([][]float64, rows)
	for i := range slices {
		slices[i] = make([]float64, cols)
		for j := range slices[i] {
			slices[i][j] = dense.At(i, j)
		}
	}
	return slices
}

// ensurePositiveDiagonal ajusta signos para diagonal positiva en R
func (m *RealMatrixAdapter) ensurePositiveDiagonal(q, r [][]float64) {
	for i := 0; i < len(r); i++ {
		if r[i][i] < 0 {
			// Ajustar columna en Q
			for j := 0; j < len(q); j++ {
				q[j][i] *= -1
			}
			// Ajustar fila en R
			for j := i; j < len(r[i]); j++ {
				r[i][j] *= -1
			}
		}
	}
}

// roundMatrix redondea valores a 'decimals' decimales
func (m *RealMatrixAdapter) roundMatrix(matrix [][]float64, decimals int) {
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = utils.RoundFloat(matrix[i][j], uint(decimals))
		}
	}
}

var ErrEmptyMatrix = errors.New("matrix cannot be empty")
