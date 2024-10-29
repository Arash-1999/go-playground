package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"

	"lets-go-book-2022/internal/config"
)

type Application struct {
	logger slog.Logger
}

var configPath = flag.String("config", "internal/config/config.yaml", "yaml config path")

func main() {
	// TODO: use signal to update flags
	config.Load(*configPath)

	app := &Application{
		logger: *config.ConfigLogger(),
	}

	// TODO: move routes to a new file (routes.go, ...)
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.homeHandler)

	mux.HandleFunc("/snippet", app.getSnippet)
	mux.HandleFunc("/snippet/create", app.postSnippet)

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", config.Configs.Server.Port),
		// TODO: pass error logger
		// ErrorLog: ,
		Handler: mux,
	}

	app.logger.Info("Starting server", "port", config.Configs.Server.Port)
	err := server.ListenAndServe()
	app.logger.Error(err.Error())
}
