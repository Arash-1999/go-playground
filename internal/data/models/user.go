package models

import (
	"context"
	"database/sql"
	"errors"
	"lets-go-book-2022/internal/lib/password"
	"strings"
	"time"

	"github.com/google/uuid"

	// "github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	CreatedAt time.Time         `json:"created_at"`
	Email     string            `json:"email"`
	FirstName string            `json:"first_name"`
	ID        string            `json:"id"`
	IsActive  bool              `json:"is_active"`
	LastName  string            `json:"last_name"`
	Password  password.Password `json:"-"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type UserModel struct {
	DB *pgxpool.Pool
}

func (m *UserModel) Insert(ctx context.Context, user *User) error {
	stmt := `INSERT INTO users (id, password_hash, email, first_name, last_name)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING id`

	id := ""

	// TODO: check this context works properly
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := m.DB.QueryRow(ctx, stmt, uuid.New(), user.Password.Hash, strings.ToLower(user.Email), user.FirstName, user.LastName).Scan(&id)

	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) Get(ctx context.Context, email string) (*User, error) {
	stmt := `SELECT id, email, first_name, last_name, password_hash, is_active, created_at, updated_at FROM users WHERE email = $1`

	user := &User{}

	// TODO: check this context works properly
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	err := m.DB.QueryRow(ctx, stmt, email).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password.Hash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsActive,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, sql.ErrNoRows
		default:
			return nil, err
		}
	}

	return user, nil
}

func (m *UserModel) Update(ctx context.Context, user *User) error {
	stmt := `UPDATE users SET frist_name = $1, last_name = $2, email = $3, password_hash = $4, is_active = $5 WHERE id = $6`

	// TODO: check this context works properly
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	args := []interface{}{
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password.Hash,
		user.IsActive,
	}
	err := m.DB.QueryRow(ctx, stmt, args...).Scan()

	if err != nil {
		// TODO: refactor error handling (create global file for errors)
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return errors.New("duplicate email")
		case errors.Is(err, sql.ErrNoRows):
			return errors.New("edit conflict")
		default:
			return err
		}
	}

	return nil
}
