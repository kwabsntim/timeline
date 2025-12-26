package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//this contains the tables for the  database

func CreateTables(db *sql.DB) error {
	wrap_table := `CREATE TABLE IF NOT EXISTS wrap(
		uuid TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		status TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := db.Exec(wrap_table)
	if err != nil {
		return err
	}
	media_table := `CREATE TABLE IF NOT EXISTS media(
		uuid TEXT PRIMARY KEY,
		wrap_uuid TEXT NOT NULL,
		filename TEXT NOT NULL,
		file_path TEXT NOT NULL,
		file_size INTEGER NOT NULL,
		mime_type TEXT NOT NULL,
		uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		photo_taken_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (wrap_uuid) REFERENCES wrap(uuid) ON DELETE CASCADE
	)`
	_, err = db.Exec(media_table)
	return err
}
