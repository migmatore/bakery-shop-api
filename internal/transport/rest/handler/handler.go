package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Deps struct {
	CustomerService CustomerService
}

type Handler struct {
	app      *fiber.App
	Customer *CustomerHandler
}

func New(deps Deps) *Handler {
	return &Handler{
		Customer: NewCustomerHandler(deps.CustomerService),
	}
}

func (h *Handler) Init() *fiber.App {
	h.app = fiber.New()

	h.app.Use(cors.New())

	api := h.app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/customer", h.Customer.GetById)

	return h.app
}
