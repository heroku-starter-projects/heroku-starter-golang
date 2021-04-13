package routes

import (
	fiber "github.com/gofiber/fiber/v2"
)

type HealthResponse struct {
	Message string `json:"message"`
}

func GetHealth(c *fiber.Ctx) error {
	healthResponse := HealthResponse{"Still alive"}

	return c.JSON(healthResponse)
}
