package base

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type Application struct {
	Logger *slog.Logger
	// TODO: add database connection pool
	Db *pgxpool.Pool
}

func InitApp(logger *slog.Logger, db *pgxpool.Pool) *Application {
	app := &Application{
		Logger: logger,
		Db:     db,
	}

	return app
}
