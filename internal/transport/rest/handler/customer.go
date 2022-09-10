package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"time"
)

type CustomerService interface {
	GetById(ctx context.Context, id int) *core.GetCustomreDTO
}

type CustomerHandler struct {
	service CustomerService
}

func NewCustomerHandler(s CustomerService) *CustomerHandler {
	return &CustomerHandler{service: s}
}

func (h *CustomerHandler) GetById(c *fiber.Ctx) error {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	defer cancel()
	customer := h.service.GetById(ctx, 2)

	return c.JSON(fiber.Map{
		"id":   customer.ID,
		"name": customer.Name,
	})
}
