package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Initialize sets up the database connection and creates necessary tables
func Initialize() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./favorites.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	createTables := `
    CREATE TABLE IF NOT EXISTS users (
        user_id TEXT PRIMARY KEY,
        username TEXT UNIQUE,
        password TEXT,
        firstname TEXT,
        lastname TEXT
    );
    CREATE TABLE IF NOT EXISTS favorites (
        user_id TEXT,
        asset_id TEXT,
        asset_type TEXT,
        description TEXT,
        PRIMARY KEY (user_id, asset_id),
        FOREIGN KEY (user_id) REFERENCES users (user_id)
    );`

	_, err = db.ExecContext(context.Background(), createTables)
	if err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return db, nil
}
