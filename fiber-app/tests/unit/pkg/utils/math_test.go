package utils_test

import (
	"fiber-app/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundFloat(t *testing.T) {
	tests := []struct {
		name      string
		input     float64
		precision uint
		expected  float64
	}{
		{"Redondeo positivo", 3.14159, 2, 3.14},
		{"Redondeo negativo", -2.71828, 3, -2.718},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, utils.RoundFloat(tt.input, tt.precision))
		})
	}
}
