package postgresql

import (
	"context"
	"database/sql"

	"github.com/begenov/tesk-kmf/internal/model"
)

type PostgreSQL struct {
	db *sql.DB
}

func NewDatabse(db *sql.DB) *PostgreSQL {
	return &PostgreSQL{db: db}
}

func (db *PostgreSQL) CreateCurrency(ctx context.Context, currency model.Currency) error {
	return nil
}

func (db *PostgreSQL) CurrencyByDateAndCode(ctx context.Context, date string, code string) ([]model.Currency, error) {
	return nil, nil
}

func (db *PostgreSQL) CurrencyByDate(ctx context.Context, date string) ([]model.Currency, error) {
	return nil, nil
}
