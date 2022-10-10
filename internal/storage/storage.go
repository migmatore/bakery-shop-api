package storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
)

type Storage struct {
	Customer *CustomerStorage
}

func New(db *pgxpool.Pool, logger *logging.Logger) *Storage {
	return &Storage{Customer: NewCustomerStorage(db, logger)}
}
