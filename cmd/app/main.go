package main

import (
	"github.com/migmatore/bakery-shop-api/internal/app"
	"github.com/migmatore/bakery-shop-api/internal/config"
	"github.com/migmatore/bakery-shop-api/pkg/logging"
	"log"
)

func main() {
	log.Print("Config initializing")
	cfg := config.GetConfig()

	log.Print("Logger initializing")
	logger := logging.GetLogger(cfg.AppConfig.LogLevel)

	app, err := app.NewApp(cfg, logger)
	if err != nil {
		log.Fatal(err)
	}

	app.Run()

}
