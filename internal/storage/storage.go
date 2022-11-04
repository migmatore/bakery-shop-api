package storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	Customer *CustomerStorage
	Product  *ProductStorage
}

func New(pool *pgxpool.Pool) *Storage {
	return &Storage{
		Customer: NewCustomerStorage(pool),
		Product:  NewProductStorage(pool),
	}
}
