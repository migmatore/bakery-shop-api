package storage

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/core"
	"github.com/migmatore/bakery-shop-api/internal/storage/psql"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type StoreStorage struct {
	pool psql.AtomicPoolClient
}

func NewStoreStorage(pool psql.AtomicPoolClient) *StoreStorage {
	return &StoreStorage{pool: pool}
}

func (s *StoreStorage) Create(ctx context.Context, store *core.CreateStore) (int, error) {
	q := `INSERT INTO stores(name, company_address_id, supplier_id) 
		  VALUES ($1, $2, $3)
          RETURNING store_id`

	var id int

	if err := s.pool.QueryRow(
		ctx,
		q,
		store.Name,
		store.StoreAddressId,
		store.SupplierId,
	).Scan(&id); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return 0, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return 0, err
	}

	return id, nil
}
