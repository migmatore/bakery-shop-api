package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/migmatore/bakery-shop-api/internal/middleware"
)

type Deps struct {
	CustomerService CustomerService
	ProductService  ProductService
	StoreService    StoreService
	EmployeeService EmployeeService
}

type Handler struct {
	app      *fiber.App
	Health   *HealthHandler
	Customer *CustomerHandler
	Product  *ProductHandler
	Store    *StoreHandler
	Employee *EmployeeHandler
}

func New(deps Deps) *Handler {
	return &Handler{
		Health:   NewHealthHandler(),
		Customer: NewCustomerHandler(deps.CustomerService),
		Product:  NewProductHandler(deps.ProductService),
		Store:    NewStoreHandler(deps.StoreService),
		Employee: NewEmployeeHandler(deps.EmployeeService),
	}
}

func (h *Handler) Init(ctx context.Context) *fiber.App {
	h.app = fiber.New()

	h.app.Use(cors.New())
	h.app.Use(logger.New())
	h.app.Use(func(c *fiber.Ctx) error {
		c.SetUserContext(ctx)

		return c.Next()
	})
	h.app.Get("/metrics", monitor.New(monitor.Config{Title: "Bakery api Metrics Page"}))
	h.app.Get("/healthz", h.Health.Health)

	api := h.app.Group("/api")
	v1 := api.Group("/v1")

	customers := v1.Group("/customers")
	customers.Post("/signin", h.Customer.Signin)
	customers.Post("/signup", h.Customer.Signup)

	customers.Get("/:id", h.Customer.GetById)
	customers.Get("/", h.Customer.GetAll)

	products := v1.Group("/products")
	products.Get("/:id", h.Product.GetOne)
	products.Get("/", h.Product.GetAll)

	products.Use(middleware.JWTProtected())
	products.Patch("/:id", h.Product.Patch)
	products.Post("/", h.Product.Create)
	products.Delete("/:id", h.Product.Delete)

	stores := v1.Group("/stores")
	stores.Post("/", h.Store.Create)

	employees := v1.Group("/employees")
	employees.Post("/signin", h.Employee.Signin)
	employees.Get("/", h.Employee.GetAll)

	return h.app
}
