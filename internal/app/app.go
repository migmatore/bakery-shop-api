package app

import (
	"log"

	"github.com/migmatore/bakery-shop-api/internal/service"
	"github.com/migmatore/bakery-shop-api/internal/storage"
	"github.com/migmatore/bakery-shop-api/internal/storage/psql"
	"github.com/migmatore/bakery-shop-api/internal/transport/rest"
	"github.com/migmatore/bakery-shop-api/internal/transport/rest/handler"
)

func Run(port string, dsn string) {
	db, err := psql.NewPostgres(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize db connection: %s", err.Error())
	}

	storages := storage.New(db)

	services := service.New(service.Deps{
		CustomerStorage: storages.Customer,
	})

	restHandlers := handler.New(handler.Deps{
		CustomerService: services.Customer,
	})

	app := restHandlers.Init()

	srv := rest.NewServer("localhost:"+port, app)
	srv.StartGracefulShutdown()
}
