package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
)

type CustomerStorage struct {
	pool   *pgxpool.Pool
	logger *logging.Logger
}

func NewCustomerStorage(pool *pgxpool.Pool, logger *logging.Logger) *CustomerStorage {
	return &CustomerStorage{pool: pool, logger: logger}
}

func (s *CustomerStorage) FindOne(ctx context.Context, id int) (*core.GetCustomreDTO, error) {
	q := `select * from customers where customers.customer_id=$1`
	var c core.GetCustomreDTO

	if err := s.pool.QueryRow(ctx, q, id).Scan(&c.ID, &c.Name); err != nil {
		s.logger.Errorf("Query error. %v", err)
		return nil, err
	}

	return &c, nil
}

func (s *CustomerStorage) FindAll(ctx context.Context) ([]*core.GetCustomreDTO, error) {
	q := `select * from customers`

	c := make([]*core.GetCustomreDTO, 0)

	rows, err := s.pool.Query(ctx, q)
	if err != nil {
		s.logger.Errorf("Query error. %v", err)
		return nil, err
	}

	for rows.Next() {
		_c := core.GetCustomreDTO{}

		err = rows.Scan(&_c.ID, &_c.Name)
		if err != nil {
			s.logger.Errorf("Query error. %v", err)
			return nil, err
		}

		c = append(c, &_c)
	}

	s.logger.Trace("Successful request")

	return c, nil
}
