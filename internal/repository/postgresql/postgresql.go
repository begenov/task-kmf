package postgresql

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/begenov/tesk-kmf/internal/model"
)

type PostgreSQL struct {
	db *sql.DB
}

func NewDatabse(db *sql.DB) *PostgreSQL {
	return &PostgreSQL{db: db}
}

func (db *PostgreSQL) CreateCurrency(ctx context.Context, rates model.Rates) error {
	tx, err := db.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	for _, currency := range rates.Currency {
		stmt := `INSERT INTO R_CURRENCY (TITLE, CODE, VALUE, A_DATE) VALUES ($1, $2, $3, $4)`
		_, err := db.db.ExecContext(ctx, stmt, currency.FullName, currency.Title, currency.Description, rates.DataTime)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (db *PostgreSQL) CurrencyByCode(ctx context.Context, date time.Time, code string) ([]model.Currency, error) {
	return nil, nil
}

func (db *PostgreSQL) CurrencyByDate(ctx context.Context, date time.Time) ([]model.Currency, error) {
	return nil, nil
}
