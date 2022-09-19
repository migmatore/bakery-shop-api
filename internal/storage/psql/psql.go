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
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			log.Printf("DB connection error. %v", err)
			return err
		}

		if err := pool.Ping(ctx); err != nil {
			log.Printf("DB ping error. %v\n", err)
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)

	//if err != nil {
	//	log.Println("Error do with tires postgresql")
	//	return nil, err
	//}

	return pool, err
}

// Auto reconnect
func Reconnect(ctx context.Context, pool *pgxpool.Pool, dsn string) {
	for {
		if err := pool.Ping(ctx); err != nil {
			//pool.Close()
			//ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
			//defer cancel()
			if pool != nil {
				p, err := pgxpool.Connect(ctx, dsn)
				if err != nil {
					log.Printf("DB reconnection error. %v", err)
					time.Sleep(1 * time.Second)
					continue
				}
				//log.Printf("pool address %p", pool)
				//log.Printf("p address %p", p)

				pool = p

				//log.Printf("pool address %p", pool)
				//log.Printf("p address %p", p)
				//
				//log.Printf("pool not nil %p, %p", pool, p)
				//
				//time.Sleep(1 * time.Second)
			}

		}

		time.Sleep(1 * time.Second)
	}
}
