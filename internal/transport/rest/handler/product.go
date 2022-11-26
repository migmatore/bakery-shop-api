package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type ProductService interface {
	GetOne(ctx context.Context, id int) (*core.Product, error)
	GetAll(ctx context.Context, queryParams map[string]string) ([]*core.Product, error)
	Patch(ctx context.Context, id int, product *core.PatchProduct) (*core.Product, error)
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

	queryParams := utils.GetQueryParams(c, "name", "price", "manufacturing_date", "expiration_date",
		"category", "manufacturer", "sort_by", "sort_order")

	products, err := h.service.GetAll(ctx, queryParams)
	if err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(products)
	//return c.Status(fiber.StatusOK).JSON(p)
}

func (h *ProductHandler) Patch(c *fiber.Ctx) error {
	ctx := c.UserContext()

	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.FiberError(c, fiber.StatusBadRequest, err)
	}

	product := new(core.PatchProduct)

	if err := c.BodyParser(product); err != nil {
		return utils.FiberError(c, fiber.StatusBadRequest, err)
	}

	newProduct, err := h.service.Patch(ctx, id, product)
	if err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(newProduct)
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	//ctx := c.UserContext()

	return nil
}
