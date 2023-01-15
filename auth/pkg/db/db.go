package db

import (
	"database/sql"
	"fmt"
)

// Create a database connection url
func CreateConnectionURL(host, username, password, database string, port int) string {
	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)
	return url
}

// Connect to Postgres database
func Connect(host, username, password, database string, port int) (*sql.DB, error) {
	db, err := sql.Open("postgres", CreateConnectionURL(host, username, password, database, 5432))
	if err != nil {
		return nil, err
	}
	return db, nil
}
