package storage

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/storage/psql"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type WishListStorage struct {
	pool psql.AtomicPoolClient
}

func NewWishListStorage(pool psql.AtomicPoolClient) *WishListStorage {
	return &WishListStorage{pool: pool}
}

func (s *WishListStorage) Create(ctx context.Context) (int, error) {
	q := `INSERT INTO wish_lists DEFAULT VALUES 
          RETURNING wish_list_id`

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
