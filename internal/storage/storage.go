package storage

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	pool     *pgxpool.Pool
	Address  *AddressStorage
	Customer *CustomerStorage
	Product  *ProductStorage
}

func New(pool *pgxpool.Pool) *Storage {
	return &Storage{
		pool:     pool,
		Address:  NewAddressStorage(pool),
		Customer: NewCustomerStorage(pool),
		Product:  NewProductStorage(pool),
	}
}

type txKey struct{}

// injectTx injects transaction to context
func injectTx(ctx context.Context, tx *pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// extractTx extracts transaction from context
func extractTx(ctx context.Context) *pgx.Tx {
	if tx, ok := ctx.Value(txKey{}).(*pgx.Tx); ok {
		return tx
	}
	return nil
}

func (s *Storage) model(ctx context.Context) *pgx.Tx {

}
