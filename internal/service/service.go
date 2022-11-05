package service

type Deps struct {
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
		Customer: NewCustomerService(deps.CustomerStorage, deps.AddressStorage),
		Product:  NewProductService(deps.ProductStorage),
	}
}
