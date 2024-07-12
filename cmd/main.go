package main

import (
	"database/sql"
	"github.com/alkosmas92/platform-go-challenge/internal/app/handlers"
	"github.com/alkosmas92/platform-go-challenge/internal/app/repository"
	"github.com/alkosmas92/platform-go-challenge/internal/app/services"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	// Set up logging
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.InfoLevel)

	// Set up the database connection
	db, err := sql.Open("sqlite3", "./favorites.db")
	if err != nil {
		logger.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	// Ensure the favorites table exists
	createTable := `
	CREATE TABLE IF NOT EXISTS favorites (
		user_id TEXT,
		asset_id TEXT,
		asset_type TEXT,
		description TEXT,
		PRIMARY KEY (user_id, asset_id)
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		logger.Fatalf("failed to create table: %v", err)
	}

	// Set up repository, service, and handlers
	favoriteRepo := repository.NewFavoriteRepository(db)
	favoriteService := services.NewFavoriteService(favoriteRepo)
	favoriteHandler := handlers.NewFavoriteHandler(favoriteService, logger)

	// Define routes and handlers
	http.HandleFunc("/favorites", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			favoriteHandler.CreateFavorite(w, r)
		case http.MethodGet:
			favoriteHandler.GetFavoritesByUserID(w, r)
		case http.MethodPut:
			favoriteHandler.UpdateFavorite(w, r)
		case http.MethodDelete:
			favoriteHandler.DeleteFavorite(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the server
	logger.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatalf("failed to start server: %v", err)
	}
}
