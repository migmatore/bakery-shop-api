package service

import "github.com/migmatore/bakery-shop-api/internal/storage"

type Deps struct {
	Transactor      storage.Transactor
	AddressStorage  AddressStorage
	CustomerStorage CustomerStorage
	ProductStorage  ProductStorage
	CartStorage     CustomerCartStorage     // TODO
	WishListStorage CustomerWishListStorage // TODO
	StoreStorage    StoreStorage
}

type Service struct {
	Address  *AddressService
	Customer *CustomerService
	Product  *ProductService
	Store    *StoreService
}

func New(deps Deps) *Service {
	return &Service{
		Address:  NewAddressService(deps.AddressStorage),
		Customer: NewCustomerService(deps.Transactor, deps.CustomerStorage, deps.AddressStorage, deps.CartStorage, deps.WishListStorage),
		Product:  NewProductService(deps.ProductStorage),
		Store:    NewStoreService(deps.Transactor, deps.StoreStorage),
	}
}
