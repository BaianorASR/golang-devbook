package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// Connect is a function that connects to the database.
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "tester:secret@tcp(localhost:3306)/test")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
