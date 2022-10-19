package service

type Deps struct {
	CustomerStorage CustomerStorage
	ProductStorage  ProductStorage
}

type Service struct {
	Customer *CustomerService
	Product  *ProductService
}

func New(deps Deps) *Service {
	return &Service{
		Customer: NewCustomerService(deps.CustomerStorage),
		Product:  NewProductService(deps.ProductStorage),
	}
}
