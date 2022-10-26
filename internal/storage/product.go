package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type ProductStorage struct {
	pool   *pgxpool.Pool
	logger *logging.Logger
}

func NewProductStorage(pool *pgxpool.Pool) *ProductStorage {
	return &ProductStorage{pool: pool}
}

func (s *ProductStorage) FindOne(ctx context.Context, id int) (*core.Product, error) {
	q := `SELECT product_id, name, description, price, manufacturing_date, expiration_date, category_id, recipe_id,
                 manufacturer_id
          FROM products WHERE product_id=$1`
	product := core.Product{}

	if err := s.pool.QueryRow(ctx, q, id).Scan(&product.ProductId, &product.Name, &product.Description, &product.Price,
		&product.ManufacturingDate, &product.ExpirationDate, &product.CategoryId, &product.RecipeId,
		&product.ManufacturerId); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return nil, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	return &product, nil
}

func (s *ProductStorage) FindAll(ctx context.Context) ([]*core.Product, error) {
	q := `SELECT product_id, name, description, price, manufacturing_date, expiration_date, category_id, recipe_id,
                 manufacturer_id
		  FROM products`

	products := make([]*core.Product, 0)

	rows, err := s.pool.Query(ctx, q)
	if err != nil {
		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		product := core.Product{}

		err := rows.Scan(&product.ProductId, &product.Name, &product.Description, &product.Price,
			&product.ManufacturingDate, &product.ExpirationDate, &product.CategoryId, &product.RecipeId,
			&product.ManufacturerId)
		if err != nil {
			logging.GetLogger(ctx).Errorf("Query error. %v", err)
			return nil, err
		}

		products = append(products, &product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
