package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenDB(driver string, dsn string) (*sql.DB, error) {
	fmt.Println(driver, dsn)
	db, err := sql.Open(driver, dsn)
	if err != nil {
		fmt.Println("ok", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
