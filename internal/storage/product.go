package storage

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/internal/storage/psql"
	"github.com/migmatore/bakery-shop-api/pkg/api/filter"
	"github.com/migmatore/bakery-shop-api/pkg/api/sort"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type ProductStorage struct {
	pool psql.AtomicPoolClient
}

func NewProductStorage(pool psql.AtomicPoolClient) *ProductStorage {
	return &ProductStorage{pool: pool}
}

func (s *ProductStorage) FindOne(ctx context.Context, id int) (*core.Product, error) {
	q := `SELECT product_id, name, image_path, description, price, manufacturing_date, expiration_date, category_id, 
       recipe_id, store_id, unit_stock, created_at, updated_at
          FROM products WHERE product_id=$1`
	product := core.Product{}

	if err := s.pool.QueryRow(ctx, q, id).Scan(
		&product.ProductId,
		&product.Name,
		&product.ImagePath,
		&product.Description,
		&product.Price,
		&product.ManufacturingDate,
		&product.ExpirationDate,
		&product.CategoryId,
		&product.RecipeId,
		&product.StoreId,
		&product.UnitStock,
		&product.CreatedAt,
		&product.UpdatedAt,
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
	q := `SELECT product_id, name, image_path, description, price, manufacturing_date, expiration_date, category_id, 
       recipe_id, store_id, unit_stock, created_at, updated_at
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
			&product.ImagePath,
			&product.Description,
			&product.Price,
			&product.ManufacturingDate,
			&product.ExpirationDate,
			&product.CategoryId,
			&product.RecipeId,
			&product.StoreId,
			&product.UnitStock,
			&product.CreatedAt,
			&product.UpdatedAt,
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
	if product.ImagePath != nil {
		updateQuery.AddUpdateColumn("image_path", product.ImagePath)
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
		updateQuery.AddUpdateColumn("recipe_id", product.RecipeId)
	}
	if product.UnitStock != nil {
		updateQuery.AddUpdateColumn("unit_stock", product.UnitStock)
	}

	updateQuery.AddUpdateColumn("updated_at", product.UpdatedAt)
	updateQuery.AddWhere("product_id", id)

	updateQuery.AddReturning("product_id", "name", "image_path", "description", "price", "manufacturing_date",
		"expiration_date", "category_id", "recipe_id", "unit_stock", "store_id", "created_at", "updated_at")

	newProduct := core.Product{}

	if err := s.pool.QueryRow(ctx, updateQuery.GetQuery(), updateQuery.GetValues()...).Scan(
		&newProduct.ProductId,
		&newProduct.Name,
		&newProduct.ImagePath,
		&newProduct.Description,
		&newProduct.Price,
		&newProduct.ManufacturingDate,
		&newProduct.ExpirationDate,
		&newProduct.CategoryId,
		&newProduct.RecipeId,
		&newProduct.UnitStock,
		&newProduct.StoreId,
		&newProduct.CreatedAt,
		&newProduct.UpdatedAt,
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

func (s *ProductStorage) Create(ctx context.Context, product *core.CreateProduct) error {
	q := `INSERT INTO products(name, image_path, description, price, manufacturing_date, expiration_date, category_id, 
                     recipe_id, store_id, unit_stock, created_at) 
		  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	if _, err := s.pool.Exec(
		ctx,
		q,
		product.Name,
		product.ImagePath,
		product.Description,
		product.Price,
		product.ManufacturingDate,
		product.ExpirationDate,
		product.CategoryId,
		product.RecipeId,
		product.StoreId,
		product.UnitStock,
		product.CreatedAt,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}
