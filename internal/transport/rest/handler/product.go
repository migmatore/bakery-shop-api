package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type ProductService interface {
	GetAll(ctx context.Context) ([]*core.Product, error)
}

type ProductHandler struct {
	service ProductService
}

func NewProductHandler(s ProductService) *ProductHandler {
	return &ProductHandler{service: s}
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
