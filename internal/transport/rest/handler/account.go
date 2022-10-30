package handler

import "github.com/gofiber/fiber/v2"

type AccountService interface {
}

type AccountHandler struct {
	service AccountService
}

func NewAccountHandler(s AccountService) *AccountHandler {
	return &AccountHandler{service: s}
}

func (h *AccountHandler) Signin(c *fiber.Ctx) error {
	return nil
}

func (h *AccountHandler) Signup(c *fiber.Ctx) error {
	return nil
}
