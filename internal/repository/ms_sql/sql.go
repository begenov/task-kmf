package sql

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/begenov/tesk-kmf/internal/model"
)

type MSSQL struct {
	db *sql.DB
}

func NewDatabse(db *sql.DB) *MSSQL {
	return &MSSQL{db: db}
}

func (db *MSSQL) CreateCurrency(ctx context.Context, rates model.Rates) error {
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

func (db *MSSQL) CurrencyByCode(ctx context.Context, date time.Time, code string) ([]model.Currency, error) {
	query := `
		SELECT ID, TITLE, CODE, VALUE, A_DATE
		FROM R_CURRENCY
		WHERE code = $1 AND a_date = $2
	`

	rows, err := db.db.QueryContext(ctx, query, code, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var currencies []model.Currency

	for rows.Next() {
		var currency model.Currency
		err = rows.Scan(&currency.ID, &currency.FullName, &currency.Title, &currency.Description, &currency.ADate)
		if err != nil {
			return nil, err
		}
		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (db *MSSQL) CurrencyByDate(ctx context.Context, date time.Time) ([]model.Currency, error) {
	query := `
		SELECT ID, TITLE, CODE, VALUE, A_DATE
		FROM R_CURRENCY
		WHERE a_date = $1
	`
	rows, err := db.db.QueryContext(ctx, query, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var currencies []model.Currency

	for rows.Next() {
		var currency model.Currency
		err = rows.Scan(&currency.ID, &currency.FullName, &currency.Title, &currency.Description, &currency.ADate)
		if err != nil {
			return nil, err
		}
		currencies = append(currencies, currency)
	}

	return currencies, nil
}
