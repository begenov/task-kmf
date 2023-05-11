package repository

import (
	"context"
	"database/sql"

	"github.com/begenov/tesk-kmf/internal/model"
	mssql "github.com/begenov/tesk-kmf/internal/repository/ms_sql"
)

type BankDB interface {
	CreateCurrency(ctx context.Context, currency model.Currency) error
}

type Repository struct {
	db BankDB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: mssql.NewDatabse(db),
	}
}
