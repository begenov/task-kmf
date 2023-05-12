package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/begenov/tesk-kmf/internal/model"
	"github.com/begenov/tesk-kmf/internal/repository/postgresql"
)

type BankDB interface {
	CreateCurrency(ctx context.Context, rates model.Rates) error
	CurrencyByCode(ctx context.Context, date time.Time, code string) ([]model.Currency, error)
	CurrencyByDate(ctx context.Context, date time.Time) ([]model.Currency, error)
}

type Repository struct {
	BankDB BankDB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		BankDB: postgresql.NewDatabse(db),
	}
}
