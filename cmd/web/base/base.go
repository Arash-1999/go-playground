package base

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"

	"lets-go-book-2022/internal/data/models"
)

type Data struct {
	Snippets *models.SnippetModel
}

type Application struct {
	Logger *slog.Logger
	Db     *Data
}

func InitApp(logger *slog.Logger, db *pgxpool.Pool) *Application {
	app := &Application{
		Logger: logger,
		Db: &Data{
			Snippets: &models.SnippetModel{DB: db},
		},
	}

	return app
}
