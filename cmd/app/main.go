package main

import (
	"context"
	"github.com/migmatore/bakery-shop-api/internal/app"
	"github.com/migmatore/bakery-shop-api/internal/config"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
)

func main() {
	//defer fgtrace.Config{Dst: fgtrace.File("fgtrace.json")}.Trace().Stop()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := logging.GetLogger(ctx)
	ctx = logging.ContextWithLogger(ctx, logger)
	logger.Info("Logger initializing")

	logger.Info("Config initializing")
	cfg := config.GetConfig(ctx)

	logger.SetLoggingLevel(cfg.AppConfig.LogLevel)

	a, err := app.NewApp(cfg)
	if err != nil {
		logger.Fatalln(err)
	}

	a.Run(ctx)

}
