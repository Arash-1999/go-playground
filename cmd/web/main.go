package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"lets-go-book-2022/internal/config"
)

type Application struct {
	logger slog.Logger
	db     *pgxpool.Pool
}

var configPath = flag.String("config", "internal/config/config.yaml", "yaml config path")

func main() {
	// TODO: use signal to update flags
	config.Load(*configPath)

	db, err := initDb(config.Configs.DB.DSN)
	logger := *config.ConfigLogger()

	if err != nil {
		logger.Error("Database Connection Error", "error", err)
	}
	app := &Application{
		logger: logger,
		db:     db,
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
	err = server.ListenAndServe()
	app.logger.Error(err.Error())
}

func initDb(dsn string) (*pgxpool.Pool, error) {
	// db, err := sql.Open("postgres", dsn)
	db, err := pgxpool.New(context.Background(), dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
