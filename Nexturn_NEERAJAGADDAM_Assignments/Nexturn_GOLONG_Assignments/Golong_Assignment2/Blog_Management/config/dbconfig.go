package db

import (
	"database/sql"

	_ "modernc.org/sqlite" // SQLite driver
)

var DB *sql.DB

// InitializeDatabase initializes the SQLite database and creates necessary tables if they don't exist
func InitializeDatabase() error {
	var err error
	DB, err = sql.Open("sqlite", "./myblogs.db")
	if err != nil {
		return err 
	}

	
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		content TEXT,
		author TEXT,
		timestamp INT
	);`)
	if err != nil {
		return err 
	}

	
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		password TEXT
	);`)
	if err != nil {
		return err 
	}

	return nil 
}


func GetDB() *sql.DB {
	return DB
}
