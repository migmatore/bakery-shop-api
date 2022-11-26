package storage

import (
	"github.com/migmatore/bakery-shop-api/internal/storage/psql"
)

type Storage struct {
	Transactor *Transact
	Address    *AddressStorage
	Customer   *CustomerStorage
	Product    *ProductStorage
	Cart       *CartStorage
	WishList   *WishListStorage
	Store      *StoreStorage
}

func New(pool psql.AtomicPoolClient) *Storage {
	return &Storage{
		Transactor: NewTransactor(pool),
		Address:    NewAddressStorage(pool),
		Customer:   NewCustomerStorage(pool),
		Product:    NewProductStorage(pool),
		Cart:       NewCartStorage(pool),
		WishList:   NewWishListStorage(pool),
		Store:      NewStoreStorage(pool),
	}
}
