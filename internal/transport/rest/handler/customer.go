package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
	"time"
)

type CustomerService interface {
	GetById(ctx context.Context, id int) (*core.Customer, error)
	GetAll(ctx context.Context) ([]*core.Customer, error)
}

type CustomerHandler struct {
	service CustomerService
}

func NewCustomerHandler(s CustomerService) *CustomerHandler {
	return &CustomerHandler{service: s}
}

// GetById TODO handle context timeout errors
func (h *CustomerHandler) GetById(c *fiber.Ctx) error {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.FiberError(c, fiber.StatusBadRequest, err)
	}

	customer, err := h.service.GetById(ctx, id)
	if err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

// GetAll TODO handle context timeout errors
func (h *CustomerHandler) GetAll(c *fiber.Ctx) error {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, 11000*time.Millisecond)
	defer cancel()

	customers, err := h.service.GetAll(ctx)
	if err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(customers)
}
