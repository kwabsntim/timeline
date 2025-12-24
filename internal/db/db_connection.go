package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "timeline.db")
	if err != nil {
		return nil, err
	}
	log.Println("database connected successfully")
	return db, nil
}
