package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./ecommerce.db")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Create the products table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		price REAL NOT NULL,
		stock INTEGER NOT NULL,
		category_id INTEGER
	);`

	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create products table: %v", err)
	}
}
