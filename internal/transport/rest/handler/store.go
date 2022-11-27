package handler

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type StoreService interface {
	Create(ctx context.Context, store *core.CreateStoreDTO) (string, error)
}

type StoreHandler struct {
	service StoreService
}

func NewStoreHandler(s StoreService) *StoreHandler {
	return &StoreHandler{service: s}
}

func (h *StoreHandler) Create(c *fiber.Ctx) error {
	ctx := c.UserContext()

	store := new(core.CreateStoreDTO)

	if err := c.BodyParser(store); err != nil {
		return utils.FiberError(c, fiber.StatusBadRequest, err)
	}

	if store.Name == "" || store.Creator.FirstName == "" || store.Creator.LastName == "" ||
		store.Creator.PhoneNumber == "" || store.Creator.Email == "" || store.Creator.Password == "" {
		return utils.FiberError(c, fiber.StatusBadRequest, errors.New("the required parameters cannot be empty"))
	}

	token, err := h.service.Create(ctx, store)
	if err != nil {
		return utils.FiberError(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
