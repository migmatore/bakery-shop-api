package service

import "github.com/migmatore/bakery-shop-api/internal/storage"

type Deps struct {
	Transactor      storage.Transactor
	AddressStorage  AddressStorage
	CustomerStorage CustomerStorage
	ProductStorage  ProductStorage
}

type Service struct {
	Address  *AddressService
	Customer *CustomerService
	Product  *ProductService
}

func New(deps Deps) *Service {
	return &Service{
		Address:  NewAddressService(deps.AddressStorage),
		Customer: NewCustomerService(deps.Transactor, deps.CustomerStorage, deps.AddressStorage),
		Product:  NewProductService(deps.ProductStorage),
	}
}
