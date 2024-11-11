package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"lets-go-book-2022/cmd/web/base"
	"lets-go-book-2022/cmd/web/base/data"
	"lets-go-book-2022/cmd/web/handlers"
	"lets-go-book-2022/internal/config"
)

var configPath = flag.String("config", "internal/config/config.yaml", "yaml config path")

func main() {
	// TODO: update config by signal
	config.Load(*configPath)

	logger := config.ConfigLogger()
	db := data.NewData()

	defer db.Close()

	app := base.InitApp(logger, db)

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", config.Configs.Server.Port),
		// TODO: pass error logger
		// ErrorLog: ,
		Handler:      handlers.Routes(app),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Info("Starting server on :", "port", config.Configs.Server.Port)
	err := server.ListenAndServe()
	log.Fatal(err)
}
