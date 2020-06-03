package postgres

import (
	"github.com/go-pg/pg/v9"
	_ "github.com/jackc/pgx/v4"
)

func New(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)
	return db
}
