package app

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/config"
	"github.com/migmatore/bakery-shop-api/internal/service"
	"github.com/migmatore/bakery-shop-api/internal/storage"
	"github.com/migmatore/bakery-shop-api/internal/storage/psql"
	"github.com/migmatore/bakery-shop-api/internal/transport/rest"
	"github.com/migmatore/bakery-shop-api/internal/transport/rest/handler"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
)

type App struct {
	cfg    *config.Config
	logger *logging.Logger
}

func NewApp(cfg *config.Config, logger *logging.Logger) (App, error) {
	return App{
		cfg: cfg, logger: logger,
	}, nil
}

func (a *App) Run() {
	a.logger.Info("Start app initializing...")

	a.logger.Info("Database connection initializing...")
	pool, err := psql.NewPostgres(context.Background(), 3, a.cfg, a.logger)
	if err != nil {
		a.logger.Fatalf("Failed to initialize db connection: %s", err.Error())
	}

	a.logger.Info("Database reconnection goroutine initializing...")
	go psql.Reconnect(context.Background(), pool, a.cfg, a.logger)

	a.logger.Info("Storages initializing...")
	storages := storage.New(pool)

	a.logger.Info("Services initializing...")
	services := service.New(service.Deps{
		CustomerStorage: storages.Customer,
	})

	a.logger.Info("Handlers initializing...")
	restHandlers := handler.New(handler.Deps{
		CustomerService: services.Customer,
	})

	app := restHandlers.Init()

	a.logger.Info("Server starting...")
	srv := rest.NewServer(a.cfg.Listen.BindIP+":"+a.cfg.Listen.Port, app, pool, a.logger)
	srv.StartWithGracefulShutdown()
}
