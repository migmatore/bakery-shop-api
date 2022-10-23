package storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	Customer *CustomerStorage
	Product  *ProductStorage
}

func New(db *pgxpool.Pool) *Storage {
	return &Storage{
		Customer: NewCustomerStorage(db),
		Product:  NewProductStorage(db),
	}
}
