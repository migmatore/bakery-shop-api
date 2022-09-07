package psql

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

func NewPostgres(ctx context.Context, maxAttempts int, dsn string) (*pgxpool.Pool, error) {
	var err error
	var pool *pgxpool.Pool

	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		log.Fatal("error do with tires postgresql")
	}

	return pool, nil
}
