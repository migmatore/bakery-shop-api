package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/pkg/api/filter"
	"github.com/migmatore/bakery-shop-api/pkg/api/sort"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
	"strings"
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

func (s *ProductStorage) Patch(ctx context.Context, id int, product *core.PatchProduct) (*core.Product, error) {
	q := `UPDATE products SET `
	qCols := make([]string, 0)
	qArgs := make([]interface{}, 0)

	if product.Name != nil {
		qCols = append(qCols, fmt.Sprintf(`name = $%d`, len(qCols)+1))
		qArgs = append(qArgs, product.Name)
	}

	if product.Price != nil {
		qCols = append(qCols, fmt.Sprintf(`price = $%d`, len(qCols)+1))
		qArgs = append(qArgs, product.Price)
	}

	q += strings.Join(qCols, ",") + fmt.Sprintf(` WHERE product_id = $%d`, len(qArgs)+1)
	qArgs = append(qArgs, id)

	q += ` RETURNING product_id, name, description, price, manufacturing_date, expiration_date, category_id, recipe_id,
           manufacturer_id`

	newProduct := core.Product{}

	if err := s.pool.QueryRow(ctx, q, qArgs...).Scan(&newProduct.ProductId, &newProduct.Name, &newProduct.Description, &newProduct.Price,
		&newProduct.ManufacturingDate, &newProduct.ExpirationDate, &newProduct.CategoryId, &newProduct.RecipeId,
		&newProduct.ManufacturerId); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return nil, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	return &newProduct, nil
}
