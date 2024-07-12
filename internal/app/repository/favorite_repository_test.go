package repository_test

import (
	"context"
	"database/sql"
	"github.com/alkosmas92/platform-go-challenge/internal/app/models"
	"github.com/alkosmas92/platform-go-challenge/internal/app/repository"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestFavoriteDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}

	createTable := `
	CREATE TABLE favorites (
		user_id TEXT,
		asset_id TEXT,
		asset_type TEXT,
		description TEXT,
		PRIMARY KEY(user_id, asset_id)
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		t.Fatalf("failed to create test table: %v", err)
	}

	return db
}

func TestFavoriteRepository(t *testing.T) {
	db := setupTestFavoriteDB(t)
	defer db.Close()

	repo := repository.NewFavoriteRepository(db)
	ctx := context.Background()

	favorite := &models.Favorite{
		UserID:      "1",
		AssetID:     "101",
		AssetType:   "book",
		Description: "A great book",
	}

	// Test CreateFavorite
	err := repo.CreateFavorite(ctx, favorite)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test GetFavoritesByUserID
	favorites, err := repo.GetFavoritesByUserID(ctx, "1", 10, 0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(favorites) != 1 || favorites[0].Description != "A great book" {
		t.Errorf("expected favorite description %v, got %v", "A great book", favorites[0].Description)
	}

	// Test UpdateFavorite
	updatedFavorite := &models.Favorite{
		UserID:      "1",
		AssetID:     "101",
		AssetType:   "book",
		Description: "An awesome book",
	}

	err = repo.UpdateFavorite(ctx, "1", "101", updatedFavorite)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	favorites, err = repo.GetFavoritesByUserID(ctx, "1", 10, 0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(favorites) != 1 || favorites[0].Description != "An awesome book" {
		t.Errorf("expected favorite description %v, got %v", "An awesome book", favorites[0].Description)
	}

	// Test DeleteFavorite
	err = repo.DeleteFavorite(ctx, "1", "101")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	favorites, err = repo.GetFavoritesByUserID(ctx, "1", 10, 0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(favorites) != 0 {
		t.Errorf("expected 0 favorites, got %v", len(favorites))
	}
}
