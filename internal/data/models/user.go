package models

import (
	"context"
	"lets-go-book-2022/internal/lib/password"
	"strings"
	"time"

	"github.com/google/uuid"

	// "github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	ID        string    `json:"id"`
	IsActive  bool      `json:"is_active"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"-"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserModel struct {
	DB *pgxpool.Pool
}

func (m *UserModel) Insert(ctx context.Context) (string, error) {
	stmt := `INSERT INTO users (id, password_hash, email, first_name, last_name)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING id;`

	id := ""
	pass := ""

	password := &password.Password{}
	err := password.Set(pass)
	if err != nil {
		return "", err
	}

	err = m.DB.QueryRow(ctx, stmt, uuid.New(), password.Hash, strings.ToLower(""), "", "").Scan(&id)

	if err != nil {
		return "", err
	}
	// 'c88329fa-f684-4cb0-912b-08ef5a4ae645', '$2y$12$9 btNEp7hzgzDM3bNyh6sY.jqqcHcGhT1JAel8sbG7ZQCQuVLef0sq', 'test-2@gmail.com', 'test-2-1', 'test-2-2'
	return "", nil
}
