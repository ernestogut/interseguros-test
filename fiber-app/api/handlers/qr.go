package handlers

import (
	"fiber-app/internal/domain/qr"
	"fiber-app/internal/infrastructure/httpclient"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type QRHandler struct {
	service    qr.QRService
	httpClient *httpclient.Client
}

func NewQRHandler(s qr.QRService, c *httpclient.Client) *QRHandler {
	return &QRHandler{service: s, httpClient: c}
}

func (h *QRHandler) ProcessQR(c *fiber.Ctx) error {
	var request struct {
		Data [][]float64 `json:"data"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request format"})
	}

	// 1. Calcular QR
	q, r, err := h.service.Factorize(request.Data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// 2. Enviar a Calcular Stats en Node
	var stats interface{}
	if h.httpClient != nil {
		token := c.Get("Authorization")
		headers := map[string]string{}
		if token != "" {
			headers["Authorization"] = token
		}
		fmt.Println("Sending stats to Node with headers:", headers)
		err := h.httpClient.PostWithHeaders("/stats", fiber.Map{"q": q, "r": r}, headers, &stats)
		if err != nil {
			fmt.Println("Error al llamar a Node:", err)
			stats = nil
		}
		fmt.Println("Received stats from Node:", stats)
	}

	return c.JSON(fiber.Map{
		"q":     q,
		"r":     r,
		"stats": stats,
	})
}
