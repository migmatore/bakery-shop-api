package storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	Customer *CustomerStorage
}

func New(db *pgxpool.Pool) *Storage {
	return &Storage{Customer: NewCustomerStorage(db)}
}
