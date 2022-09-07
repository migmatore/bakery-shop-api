package storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type CustomerStorage struct {
	db *pgxpool.Pool
}

func NewCustomerStorage(db *pgxpool.Pool) *CustomerStorage {
	return &CustomerStorage{db: db}
}
