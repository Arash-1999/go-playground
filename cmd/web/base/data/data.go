package data

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"lets-go-book-2022/internal/config"
)

func NewData() *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), config.Configs.Database.Dsn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	return db
}
