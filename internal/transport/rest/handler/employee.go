package handler

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type EmployeeService interface {
	Signin(ctx context.Context, employeeAcc *core.SigninEmployeeDTO) (*core.EmployeeTokenMetadata, error)
	GetAll(ctx context.Context) ([]*core.Employee, error)
}

type EmployeeHandler struct {
	service EmployeeService
}

func NewEmployeeHandler(s EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: s}
}

func (h *EmployeeHandler) Signin(c *fiber.Ctx) error {
	ctx := c.UserContext()
	employeeAcc := new(core.SigninEmployeeDTO)

	if err := c.BodyParser(employeeAcc); err != nil {
		return utils.FiberError(c, fiber.StatusBadRequest, err)
	}

	if employeeAcc.Email == "" || employeeAcc.Password == "" {
		return utils.FiberError(c, fiber.StatusBadRequest, errors.New("the required parameters cannot be empty"))
	}

	token, err := h.service.Signin(ctx, employeeAcc)
	if err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(token)
}

func (h *EmployeeHandler) GetAll(c *fiber.Ctx) error {
	ctx := c.UserContext()

	employees, err := h.service.GetAll(ctx)
	if err != nil {
		return err
	}

	return c.JSON(employees)
}
