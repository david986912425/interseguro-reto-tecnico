package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go-backend/internal/services"
)

func PostMatrix(c *fiber.Ctx) error {
	var input [][]float64

	if err := json.Unmarshal(c.Body(), &input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Formato inválido: se esperaba una matriz de números",
		})
	}

	token := c.Get("Authorization")
	result, err := services.ProcessMatrix(input, token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}
