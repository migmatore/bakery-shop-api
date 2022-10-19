package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
)

type ProductStorage struct {
	pool   *pgxpool.Pool
	logger *logging.Logger
}

func NewProductStorage(pool *pgxpool.Pool, logger *logging.Logger) *ProductStorage {
	return &ProductStorage{pool: pool, logger: logger}
}

func (s *ProductStorage) FindAll(ctx context.Context) ([]*core.Product, error) {
	q := `select * from products`

	products := make([]*core.Product, 0)

	rows, err := s.pool.Query(ctx, q)
	if err != nil {
		s.logger.Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		product := core.Product{}

		err := rows.Scan(&product.ProductId, &product.Name, &product.Description, &product.Price,
			&product.ManufacturingDate, &product.ExpirationDate, &product.CategoryId, &product.RecipeId,
			&product.ManufacturerId)
		if err != nil {
			s.logger.Errorf("Query error. %v", err)
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}
