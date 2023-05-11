package sql

import (
	"context"
	"database/sql"

	"github.com/begenov/tesk-kmf/internal/model"
)

type Database struct {
	db *sql.DB
}

func NewDatabse(db *sql.DB) *Database {
	return &Database{db: db}
}

func (db *Database) CreateCurrency(ctx context.Context, currency model.Currency) error {
	return nil
}
