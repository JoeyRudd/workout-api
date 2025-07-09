package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewConnection() (*sql.DB, error) {
	connectionStr := "host=localhost port=5432 user=joey password=postgres dbname=workout_db sslmode=disable"

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
