package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//this contains the tables for the  database

func CreateTables(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS wrap(
		uuid TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		status TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := db.Exec(query)
	return err
}
