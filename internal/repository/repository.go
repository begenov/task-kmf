package repository

import (
	"context"
	"database/sql"

	"github.com/begenov/tesk-kmf/internal/model"
	"github.com/begenov/tesk-kmf/internal/repository/postgresql"
)

type BankDB interface {
	CreateCurrency(ctx context.Context, rates model.Rates) error
	CurrencyByDateAndCode(ctx context.Context, date string, code string) ([]model.Currency, error)
	CurrencyByDate(ctx context.Context, date string) ([]model.Currency, error)
}

type Repository struct {
	BankDB BankDB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		BankDB: postgresql.NewDatabse(db),
	}
}
