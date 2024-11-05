package models

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Snippet struct {
	ID      int
	Title   string
	Content pgtype.Text
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *pgxpool.Pool
}

func (m *SnippetModel) Insert(ctx context.Context, title string, content string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
  VALUES ($1, $2, now(), now() + INTERVAL '365 DAY')
  RETURNING id`

	id := 0
	err := m.DB.QueryRow(ctx, stmt, title, content).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(ctx context.Context, id int) (*Snippet, error) {
	stmt := ` SELECT id, title, content, expires FROM snippets
  WHERE expires > now() AND id = $1`
	s := &Snippet{}

	err := m.DB.QueryRow(ctx, stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Expires)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest(ctx context.Context) ([]*Snippet, error) {
	// TODO: use LIMIT and OFFSET for pagination
	stmt := `SELECT id, title, content, created, expires FROM snippets
  WHERE expires > now() ORDER BY id DESC LIMIT 10 OFFSET 0`

	rows, err := m.DB.Query(ctx, stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	snippets := []*Snippet{}

	for rows.Next() {
		s := &Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
