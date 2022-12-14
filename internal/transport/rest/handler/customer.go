package handler

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
	"time"
)

type CustomerService interface {
	GetById(ctx context.Context, id int) (*core.Customer, error)
	GetAll(ctx context.Context) ([]*core.Customer, error)
	Signup(ctx context.Context, customer *core.CreateCustomerDTO) (*core.CustomerTokenMetadata, error)
	Signin(ctx context.Context, customer *core.SigninCustomerDTO) (*core.CustomerTokenMetadata, error)
}

type CustomerHandler struct {
	service CustomerService
}

func NewCustomerHandler(s CustomerService) *CustomerHandler {
	return &CustomerHandler{service: s}
}

func (h *CustomerHandler) Signin(c *fiber.Ctx) error {
	ctx := c.UserContext()
	customerAcc := new(core.SigninCustomerDTO)

	if err := c.BodyParser(customerAcc); err != nil {
		return utils.FiberError(c, fiber.StatusBadRequest, err)
	}

	if customerAcc.Email == "" || customerAcc.Password == "" {
		return utils.FiberError(c, fiber.StatusBadRequest, errors.New("the required parameters cannot be empty"))
	}

	token, err := h.service.Signin(ctx, customerAcc)
	if err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(token)
}

func (h *CustomerHandler) Signup(c *fiber.Ctx) error {
	ctx := c.UserContext()

	customer := new(core.CreateCustomerDTO)

	if err := c.BodyParser(customer); err != nil {
		return utils.FiberError(c, fiber.StatusBadRequest, err)
	}

	if customer.FirstName == "" || customer.LastName == "" || customer.ImagePath == "" || customer.PhoneNumber == "" ||
		customer.Email == "" || customer.Password == "" {
		return utils.FiberError(c, fiber.StatusBadRequest, errors.New("the required parameters cannot be empty"))
	}

	token, err := h.service.Signup(ctx, customer)
	if err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(token)
}

// GetById TODO handle context timeout errors
func (h *CustomerHandler) GetById(c *fiber.Ctx) error {
	ctx := c.UserContext()
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
	ctx := c.UserContext()
	//ctx, cancel := context.WithTimeout(ctx, 11000*time.Millisecond)
	//defer cancel()

	customers, err := h.service.GetAll(ctx)
	if err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(customers)
}
