package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func InitDB() (*sql.DB, error) {
	// Get DB path from environment
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./data/yearwrap.db"
	}

	// Ensure directory exists
	dir := filepath.Dir(dbPath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Printf("Failed to create directory %s: %v", dir, err)
		return nil, err
	}

	// Open database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Enable foreign keys
	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, err
	}

	log.Printf("✅ Database connected successfully at %s", dbPath)
	return db, nil
}
