package repository

import (
	"database/sql"

	mssql "github.com/begenov/tesk-kmf/internal/repository/ms_sql"
)

type BankDB interface {
}

type Repository struct {
	db BankDB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: mssql.NewDatabse(db),
	}
}
