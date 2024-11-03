package data

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"

	"lets-go-book-2022/internal/config"
)

func NewData(config config.Config) *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), config.Database.Dsn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	return db
}
