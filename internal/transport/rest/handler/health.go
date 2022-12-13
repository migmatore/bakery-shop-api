package handler

import (
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK)

	return nil
}
