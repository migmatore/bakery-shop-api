package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"time"
)

type CustomerService interface {
	GetById(ctx context.Context, id int) *core.GetCustomreDTO
	GetAll(ctx context.Context) []*core.GetCustomreDTO
}

type CustomerHandler struct {
	service CustomerService
}

func NewCustomerHandler(s CustomerService) *CustomerHandler {
	return &CustomerHandler{service: s}
}

func (h *CustomerHandler) GetById(c *fiber.Ctx) error {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	customer := h.service.GetById(ctx, id)

	return c.JSON(fiber.Map{
		"id":   customer.ID,
		"name": customer.Name,
	})
}

func (h *CustomerHandler) GetAll(c *fiber.Ctx) error {
	customers := h.service.GetAll(context.TODO())

	return c.JSON(customers)
}
