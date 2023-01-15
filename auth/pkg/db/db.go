package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Create a database connection url
func CreateConnectionURL(host, username, password, database string, port int) string {
	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)
	return url
}

// Connect to Postgres database
func Connect(host, username, password, database string, port int) (*sql.DB, error) {
	db, err := sql.Open("postgres", CreateConnectionURL(host, username, password, database, port))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Close the database connection
func Close(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}

// Ping the database connection
func Ping(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}
