package storage

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/migmatore/bakery-shop-api/internal/storage/psql"
)

type Transactor interface {
	WithinTransaction(ctx context.Context, txFunc func(txCtx context.Context) error) error
}

type Transact struct {
	pool psql.AtomicPoolClient
}

func NewTransactor(pool psql.AtomicPoolClient) *Transact {
	return &Transact{pool: pool}
}

func (s *Transact) WithinTransaction(ctx context.Context, txFunc func(context context.Context) error) error {
	return s.pool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return txFunc(psql.InjectTx(ctx, tx))
	})
}
