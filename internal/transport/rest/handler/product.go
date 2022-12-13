package handler

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/jwt"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
	"time"
)

type ProductService interface {
	GetOne(ctx context.Context, id int) (*core.Product, error)
	GetAll(ctx context.Context, queryParams map[string]string) (*core.ProductPage, error)
	Patch(ctx context.Context, id int, product *core.PatchProductDTO) (*core.Product, error)
	Create(ctx context.Context, product *core.CreateProductDTO, employeeId int, storeId int) error
	Delete(ctx context.Context, id int) error
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
		"category", "store", "sort_by", "sort_order", "page", "per_page")

	productPage, err := h.service.GetAll(ctx, queryParams)
	if err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(productPage)
	//return c.Status(fiber.StatusOK).JSON(p)
}

func (h *ProductHandler) Patch(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	if now > claims.Expires || claims.Customer == true || claims.Admin == false {
		return utils.FiberError(
			c,
			fiber.StatusUnauthorized,
			errors.New("unauthorized, check expiration time or access level of your token"),
		)
	}

	ctx := c.UserContext()

	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.FiberError(c, fiber.StatusBadRequest, err)
	}

	product := new(core.PatchProductDTO)

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
	now := time.Now().Unix()

	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	if now > claims.Expires || claims.Customer == true || claims.Admin == false {
		return utils.FiberError(
			c,
			fiber.StatusUnauthorized,
			errors.New("unauthorized, check expiration time or access level of your token"),
		)
	}

	ctx := c.UserContext()
	product := new(core.CreateProductDTO)

	if err := c.BodyParser(product); err != nil {
		return utils.FiberError(c, fiber.StatusBadRequest, err)
	}

	if product.Name == "" || product.ImagePath == "" || product.Price == 0 || product.ManufacturingDate == "" ||
		product.ExpirationDate == "" || product.UnitStock < 0 {
		return utils.FiberError(c, fiber.StatusBadRequest, errors.New("the required parameters cannot be empty"))
	}

	if err := h.service.Create(ctx, product, claims.Id, claims.StoreId); err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(200).JSON(fiber.Map{
		"msg": "Product was created",
	})
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	if now > claims.Expires || claims.Customer == true || claims.Admin == false {
		return utils.FiberError(
			c,
			fiber.StatusUnauthorized,
			errors.New("unauthorized, check expiration time or access level of your token"),
		)
	}

	ctx := c.UserContext()

	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.FiberError(c, fiber.StatusBadRequest, err)
	}

	if err := h.service.Delete(ctx, id); err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "product was successfully deleted",
	})
}
