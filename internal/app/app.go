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
	cfg *config.Config
}

func NewApp(cfg *config.Config) (App, error) {
	return App{
		cfg: cfg,
	}, nil
}

func (a *App) Run(ctx context.Context) {
	logging.GetLogger(ctx).Info("Start app initializing...")

	logging.GetLogger(ctx).Info("Database connection initializing...")
	pool, err := psql.NewPostgres(ctx, 3, a.cfg)
	if err != nil {
		logging.GetLogger(ctx).Fatalf("Failed to initialize db connection: %s", err.Error())
	}

	logging.GetLogger(ctx).Info("Database reconnection goroutine initializing...")
	go psql.Reconnect(ctx, pool, a.cfg)

	logging.GetLogger(ctx).Info("Storages initializing...")
	storages := storage.New(pool)

	logging.GetLogger(ctx).Info("Services initializing...")
	services := service.New(service.Deps{
		CustomerStorage: storages.Customer,
		ProductStorage:  storages.Product,
	})

	logging.GetLogger(ctx).Info("Handlers initializing...")
	restHandlers := handler.New(handler.Deps{
		CustomerService: services.Customer,
		ProductService:  services.Product,
	})

	app := restHandlers.Init(ctx)

	logging.GetLogger(ctx).Info("Server starting...")
	srv := rest.NewServer(a.cfg.Listen.BindIP+":"+a.cfg.Listen.Port, app, pool)
	srv.StartWithGracefulShutdown(ctx)
}
