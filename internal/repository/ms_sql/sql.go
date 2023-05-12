package sql

import (
	"context"
	"database/sql"

	"github.com/begenov/tesk-kmf/internal/model"
)

type MSSQL struct {
	db *sql.DB
}

func NewDatabse(db *sql.DB) *MSSQL {
	return &MSSQL{db: db}
}

func (db *MSSQL) CreateCurrency(ctx context.Context, currency model.Currency) error {
	return nil
}

func (db *MSSQL) CurrencyByDateAndCode(ctx context.Context, date string, code string) ([]model.Currency, error) {
	return nil, nil
}

func (db *MSSQL) CurrencyByDate(ctx context.Context, date string) ([]model.Currency, error) {
	return nil, nil
}
