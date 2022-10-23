package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type CustomerStorage struct {
	pool *pgxpool.Pool
}

func NewCustomerStorage(pool *pgxpool.Pool) *CustomerStorage {
	return &CustomerStorage{pool: pool}
}

func (s *CustomerStorage) FindOne(ctx context.Context, id int) (*core.Customer, error) {
	q := `select customer_id, first_name, telephone_number from customers where customers.customer_id=$1`
	var c core.Customer

	if err := s.pool.QueryRow(ctx, q, id).Scan(&c.CustomerId, &c.FirstName, &c.TelephoneNumber); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return nil, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	return &c, nil
}

func (s *CustomerStorage) FindAll(ctx context.Context) ([]*core.Customer, error) {
	q := `select customer_id, first_name, last_name, telephone_number from customers`

	customers := make([]*core.Customer, 0)

	rows, err := s.pool.Query(ctx, q)
	if err != nil {
		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		customer := core.Customer{}

		err := rows.Scan(&customer.CustomerId, &customer.FirstName, &customer.LastName, &customer.TelephoneNumber)
		if err != nil {
			logging.GetLogger(ctx).Errorf("Query error. %v", err)
			return nil, err
		}

		customers = append(customers, &customer)
	}

	return customers, nil
}
