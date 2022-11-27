package service

import "github.com/migmatore/bakery-shop-api/internal/storage"

type Deps struct {
	Transactor      storage.Transactor
	AddressStorage  AddressStorage
	CustomerStorage CustomerStorage
	ProductStorage  ProductStorage
	CartStorage     CustomerCartStorage     // TODO interface from service
	WishListStorage CustomerWishListStorage // TODO interface from service
	StoreStorage    StoreStorage            // TODO interface from service
	EmployeeStorage EmployeeStorage
}

type Service struct {
	Address  *AddressService
	Customer *CustomerService
	Product  *ProductService
	Store    *StoreService
	Employee *EmployeeService
}

func New(deps Deps) *Service {
	return &Service{
		Address:  NewAddressService(deps.AddressStorage),
		Customer: NewCustomerService(deps.Transactor, deps.CustomerStorage, deps.AddressStorage, deps.CartStorage, deps.WishListStorage),
		Product:  NewProductService(deps.ProductStorage),
		Store:    NewStoreService(deps.Transactor, deps.StoreStorage, deps.EmployeeStorage),
		Employee: NewEmployeeService(deps.EmployeeStorage),
	}
}
