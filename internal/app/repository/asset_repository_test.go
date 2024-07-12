package repository_test

import (
	"context"
	"database/sql"
	"github.com/alkosmas92/platform-go-challenge/internal/app/models"
	"github.com/alkosmas92/platform-go-challenge/internal/app/repository"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}

	createTable := `
    CREATE TABLE assets (
        id TEXT PRIMARY KEY,
        type TEXT,
        description TEXT,
        audience BLOB,
        chart BLOB,
        insight BLOB
    );`

	_, err = db.Exec(createTable)
	if err != nil {
		t.Fatalf("failed to create test table: %v", err)
	}

	return db
}

func TestAssetRepository(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repository.NewAssetRepository(db)
	ctx := context.Background()

	audience := &models.Audience{
		Gender:             "Male",
		BirthCountry:       "USA",
		AgeGroup:           "18-24",
		HoursOnSocialMedia: 2,
		PurchasesLastMonth: 5,
	}

	asset := &models.Asset{
		ID:          "1",
		Type:        "audience",
		Description: "Test audience asset",
		Audience:    audience,
	}

	// Test CreateAsset
	err := repo.CreateAsset(ctx, asset)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test GetAssetByID
	fetchedAsset, err := repo.GetAssetByID(ctx, "1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if fetchedAsset.Audience == nil || fetchedAsset.Audience.Gender != "Male" {
		t.Errorf("expected audience gender %v, got %v", "Male", fetchedAsset.Audience.Gender)
	}

	// Test UpdateAsset
	updatedAudience := &models.Audience{
		Gender:             "Female",
		BirthCountry:       "Canada",
		AgeGroup:           "25-34",
		HoursOnSocialMedia: 3,
		PurchasesLastMonth: 6,
	}

	updatedAsset := &models.Asset{
		ID:          "1",
		Type:        "audience",
		Description: "Updated audience asset",
		Audience:    updatedAudience,
	}

	err = repo.UpdateAssetByID(ctx, "1", updatedAsset)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	fetchedAsset, err = repo.GetAssetByID(ctx, "1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if fetchedAsset.Audience == nil || fetchedAsset.Audience.Gender != "Female" {
		t.Errorf("expected audience gender %v, got %v", "Female", fetchedAsset.Audience.Gender)
	}

	// Test DeleteAsset
	err = repo.DeleteAssetByID(ctx, "1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = repo.GetAssetByID(ctx, "1")
	if err == nil {
		t.Fatalf("expected error, got none")
	}
}
