package integration_test

import (
	"encoding/json"
	"fiber-app/pkg/server"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIIntegration(t *testing.T) {
	os.Setenv("JWT_SECRET", "IntersegurosSecretKey")
	app := server.NewTestApp()
	t.Log(os.Getenv("JWT_SECRET"))
	fmt.Println(os.Getenv("JWT_SECRET"))
	t.Run("Full QR processing flow", func(t *testing.T) {
		body := `{"data":[[1,2],[3,4]]}`

		req := httptest.NewRequest("POST", "/process", strings.NewReader(body))
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIn0.P955zJ1Tdo9WfxqbW0SV_lOFtAYiYC1gFXllqPnVdHQ")
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// Verificar respuesta
		var result struct {
			Q [][]float64 `json:"q"`
			R [][]float64 `json:"r"`
		}
		err = json.NewDecoder(resp.Body).Decode(&result)
		assert.NoError(t, err)
		assert.NotEmpty(t, result.Q)
		assert.NotEmpty(t, result.R)

	})
}
