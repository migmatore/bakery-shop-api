package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"time"
)

type CustomerService interface {
	GetById(ctx context.Context, id int) (*core.GetCustomreDTO, error)
	GetAll(ctx context.Context) ([]*core.GetCustomreDTO, error)
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

	customer, err := h.service.GetById(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"id":   customer.ID,
		"name": customer.Name,
	})
}

func (h *CustomerHandler) GetAll(c *fiber.Ctx) error {
	customers, err := h.service.GetAll(context.TODO())
	if err != nil {
		return err
	}

	return c.JSON(customers)
}
