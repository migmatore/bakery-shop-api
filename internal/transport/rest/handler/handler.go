package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
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

	//h.app.Use(cors.New())
	h.app.Use(logger.New())
	h.app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	api := h.app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/customer/:id", h.Customer.GetById)
	v1.Get("/customers", h.Customer.GetAll)

	return h.app
}
