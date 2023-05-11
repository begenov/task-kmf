package sql

import "database/sql"

type Database struct {
	db *sql.DB
}

func NewDatabse(db *sql.DB) *Database {
	return &Database{db: db}
}
