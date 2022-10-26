package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type ProductService interface {
	GetOne(ctx context.Context, id int) (*core.Product, error)
	GetAll(ctx context.Context) ([]*core.Product, error)
}

type ProductHandler struct {
	service ProductService
}

func NewProductHandler(s ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) GetOne(c *fiber.Ctx) error {
	ctx := c.UserContext()

	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.FiberError(c, fiber.StatusBadRequest, err)
	}

	product, err := h.service.GetOne(ctx, id)
	if err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	ctx := c.UserContext()
	//_ctx, cancel := context.WithTimeout(ctx, 1*time.Millisecond)
	//defer cancel()
	products, err := h.service.GetAll(ctx)
	if err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(products)
}
