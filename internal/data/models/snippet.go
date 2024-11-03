package models

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *pgxpool.Pool
}

func (m *SnippetModel) Insert() (int, error) {
	return 0, nil
}

func (m *SnippetModel) Get() (*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() (*Snippet, error) {
	return nil, nil
}
