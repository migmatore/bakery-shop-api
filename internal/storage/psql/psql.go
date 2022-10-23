package psql

import (
	"context"
	"fmt"
	"github.com/migmatore/bakery-shop-api/internal/config"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

func NewPostgres(ctx context.Context, maxAttempts int, cfg *config.Config) (*pgxpool.Pool, error) {
	var err error
	var pool *pgxpool.Pool

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?pool_max_conns=%s",
		cfg.DBConnection.Username,
		cfg.DBConnection.Password,
		cfg.DBConnection.Host,
		cfg.DBConnection.Port,
		cfg.DBConnection.DB,
		cfg.DBConnection.MaxConns,
	)

	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			logging.GetLogger(ctx).Errorf("DB connection error. %v", err)
			return err
		}

		if err := pool.Ping(ctx); err != nil {
			logging.GetLogger(ctx).Errorf("DB ping error. %v\n", err)
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)

	return pool, err
}

// Reconnect Auto reconnecting to db
func Reconnect(ctx context.Context, pool *pgxpool.Pool, cfg *config.Config) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?pool_max_conns=%s",
		cfg.DBConnection.Username,
		cfg.DBConnection.Password,
		cfg.DBConnection.Host,
		cfg.DBConnection.Port,
		cfg.DBConnection.DB,
		cfg.DBConnection.MaxConns,
	)

	for {
		if err := pool.Ping(ctx); err != nil {
			pool.Close()
			//ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
			//defer cancel()
			if pool != nil {
				p, err := pgxpool.Connect(ctx, dsn)
				if err != nil {
					logging.GetLogger(ctx).Errorf("DB reconnection error. %v", err)
					time.Sleep(1 * time.Second)

					continue
				}

				pool = p
			}

		}

		time.Sleep(1 * time.Second)
	}
}
