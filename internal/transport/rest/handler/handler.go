package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type Deps struct {
	CustomerService CustomerService
	ProductService  ProductService
}

type Handler struct {
	app      *fiber.App
	Customer *CustomerHandler
	Product  *ProductHandler
}

func New(deps Deps) *Handler {
	return &Handler{
		Customer: NewCustomerHandler(deps.CustomerService),
		Product:  NewProductHandler(deps.ProductService),
	}
}

func (h *Handler) Init(ctx context.Context) *fiber.App {
	h.app = fiber.New()

	//h.app.Use(cors.New())
	//h.app.Use(logger.New())
	h.app.Use(func(c *fiber.Ctx) error {
		c.SetUserContext(ctx)

		return c.Next()
	})
	h.app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	api := h.app.Group("/api")
	v1 := api.Group("/v1")

	customer := v1.Group("/customers")

	customer.Post("/signin", h.Customer.Signin)
	customer.Post("/signup", h.Customer.Signup)

	customer.Get("/:id", h.Customer.GetById)
	customer.Get("/", h.Customer.GetAll)

	product := v1.Group("/product")

	product.Get("/:id", h.Product.GetOne)
	product.Get("/all", h.Product.GetAll)
	product.Patch("/:id", h.Product.Patch)

	return h.app
}
