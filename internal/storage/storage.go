package storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	Address  *AddressStorage
	Customer *CustomerStorage
	Product  *ProductStorage
}

func New(pool *pgxpool.Pool) *Storage {
	return &Storage{
		Address:  NewAddressStorage(pool),
		Customer: NewCustomerStorage(pool),
		Product:  NewProductStorage(pool),
	}
}
