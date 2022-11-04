package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/internal/storage/psql"
	"github.com/migmatore/bakery-shop-api/pkg/api/filter"
	"github.com/migmatore/bakery-shop-api/pkg/api/sort"
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

	if err := s.pool.QueryRow(ctx, q, id).Scan(
		&product.ProductId,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.ManufacturingDate,
		&product.ExpirationDate,
		&product.CategoryId,
		&product.RecipeId,
		&product.ManufacturerId,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return nil, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	return &product, nil
}

func (s *ProductStorage) FindAll(ctx context.Context, filterOptions []filter.Option, sortOption sort.Option) ([]*core.Product, error) {
	q := `SELECT product_id, name, description, price, manufacturing_date, expiration_date, category_id, recipe_id,
                 manufacturer_id
		  FROM products`

	q = filter.EnrichQueryWithFilter(q, filterOptions)
	q = sort.EnrichQueryWithSort(q, sortOption)

	products := make([]*core.Product, 0)

	rows, err := s.pool.Query(ctx, q)
	if err != nil {
		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		product := core.Product{}

		err := rows.Scan(
			&product.ProductId,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.ManufacturingDate,
			&product.ExpirationDate,
			&product.CategoryId,
			&product.RecipeId,
			&product.ManufacturerId,
		)
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

func (s *ProductStorage) Patch(ctx context.Context, id int, product *core.PatchProduct) (*core.Product, error) {
	updateQuery := psql.NewSQLUpdateBuilder("products")

	if product.Name != nil {
		updateQuery.AddUpdateColumn("name", product.Name)
	}
	if product.Description != nil {
		updateQuery.AddUpdateColumn("description", product.Description)
	}
	if product.Price != nil {
		updateQuery.AddUpdateColumn("price", product.Price)
	}
	if product.ManufacturingDate != nil {
		updateQuery.AddUpdateColumn("manufacturing_date", product.ManufacturingDate)
	}
	if product.ExpirationDate != nil {
		updateQuery.AddUpdateColumn("expiration_date", product.ExpirationDate)
	}
	if product.CategoryId != nil {
		updateQuery.AddUpdateColumn("category_id", product.CategoryId)
	}
	if product.RecipeId != nil {
		updateQuery.AddUpdateColumn("price", product.RecipeId)
	}
	// TODO set current manufacturer id
	if product.ManufacturerId != nil {
		updateQuery.AddUpdateColumn("manufacturer_id", product.ManufacturerId)
	}

	updateQuery.AddWhere("product_id", id)

	updateQuery.AddReturning("product_id", "name", "description", "price", "manufacturing_date", "expiration_date",
		"category_id", "recipe_id", "manufacturer_id")

	newProduct := core.Product{}

	if err := s.pool.QueryRow(ctx, updateQuery.GetQuery(), updateQuery.GetValues()...).Scan(
		&newProduct.ProductId,
		&newProduct.Name,
		&newProduct.Description,
		&newProduct.Price,
		&newProduct.ManufacturingDate,
		&newProduct.ExpirationDate,
		&newProduct.CategoryId,
		&newProduct.RecipeId,
		&newProduct.ManufacturerId,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return nil, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	return &newProduct, nil
}
