package app

import (
	"context"
	"log"

	"github.com/migmatore/bakery-shop-api/internal/service"
	"github.com/migmatore/bakery-shop-api/internal/storage"
	"github.com/migmatore/bakery-shop-api/internal/storage/psql"
	"github.com/migmatore/bakery-shop-api/internal/transport/rest"
	"github.com/migmatore/bakery-shop-api/internal/transport/rest/handler"
)

func Run(port string, dsn string) {
	pool, err := psql.NewPostgres(context.Background(), 3, dsn)
	if err != nil {
		log.Printf("Failed to initialize db connection: %s", err.Error())
	}

	//log.Printf("pool address %p", pool)

	go psql.Reconnect(context.Background(), pool, dsn)

	storages := storage.New(pool)

	services := service.New(service.Deps{
		CustomerStorage: storages.Customer,
	})

	restHandlers := handler.New(handler.Deps{
		CustomerService: services.Customer,
	})

	app := restHandlers.Init()

	srv := rest.NewServer("localhost:"+port, app, pool)
	srv.StartGracefulShutdown()
}
