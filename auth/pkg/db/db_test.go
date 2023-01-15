package db

import (
	"database/sql"
	"testing"
)

// Test the database connection string
func TestCreateConnectionURL(t *testing.T) {
	url := CreateConnectionURL("localhost", "postgres", "postgres", "postgres", 5432)
	if url != "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable" {
		t.Error("Failed to create a valid connection string")
	}
}

// Test the database connection
func TestConnect(t *testing.T) {
	conn, err := Connect("localhost", "postgres", "postgres", "postgres", 5432)
	if err != nil {
		t.Error("Failed to connect to the database")
	}
	defer conn.Close()
}

// Test the database connection
func TestConnectWithInvalidCredentials() (*sql.DB, error) {
	return Connect("localhost", "postgres", "invalid", "postgres", 5432)
}

// Test the database connection
func TestConnectWithInvalidDatabase() (*sql.DB, error) {
	return Connect("localhost", "postgres", "postgres", "invalid", 5432)
}

// Test the database connection
func TestConnectWithInvalidHost() (*sql.DB, error) {
	return Connect("invalid", "postgres", "postgres", "postgres", 5432)
}

// Test the database connection
func TestConnectWithInvalidPort() (*sql.DB, error) {
	return Connect("localhost", "postgres", "postgres", "postgres", 0)
}

// Test the database connection
func TestConnectWithInvalidUsername() (*sql.DB, error) {
	return Connect("localhost", "invalid", "postgres", "postgres", 5432)
}

// Test the database connection
func TestConnectWithInvalidURL() (*sql.DB, error) {
	return Connect("localhost", "postgres", "postgres", "postgres", 5432)
}
