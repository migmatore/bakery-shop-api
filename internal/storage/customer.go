package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"log"
)

type CustomerStorage struct {
	pool *pgxpool.Pool
}

func NewCustomerStorage(pool *pgxpool.Pool) *CustomerStorage {
	return &CustomerStorage{pool: pool}
}

func (s *CustomerStorage) FindOne(ctx context.Context, id int) *core.GetCustomreDTO {
	q := `select * from customers where customers.id=$1`

	var c core.GetCustomreDTO

	if err := s.pool.QueryRow(ctx, q, id).Scan(&c.ID, &c.Name); err != nil {
		log.Printf("Query error. %v", err)
	}

	return &c
}
