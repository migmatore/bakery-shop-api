package storage

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/storage/psql"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type CartStorage struct {
	pool psql.AtomicPoolClient
}

func NewCartStorage(pool psql.AtomicPoolClient) *CartStorage {
	return &CartStorage{pool: pool}
}

func (s *CartStorage) Create(ctx context.Context) (int, error) {
	q := `INSERT INTO carts DEFAULT VALUES 
          RETURNING cart_id`

	var id int

	if err := s.pool.QueryRow(ctx, q).Scan(&id); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return 0, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return 0, err
	}

	return id, nil
}
