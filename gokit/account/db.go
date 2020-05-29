package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

var ErrDB = errors.New("Unable to handle DB request")

type DB struct {
	db     *sql.DB
	logger log.Logger
}

func NewDB(db *sql.DB, logger log.Logger) Database {
	return &DB{
		db:     db,
		logger: log.With(logger, "db", "sql"),
	}
}

func (DB *DB) CreateUser(ctx context.Context, user User) error {
	sql := `
	INSERT INTO users (id, email, password)
	VALUES ($1, $2, $3)
	`

	if user.Email == "" || user.Password == "" {
		return ErrDB
	}

	if _, err := DB.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password); err != nil {
		return err
	}

	return nil
}

func (DB *DB) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	err := DB.db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email)
	if err != nil {
		return "", ErrDB
	}

	return email, nil
}
