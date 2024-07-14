package server

import (
	"database/sql"
	"net/http"

	"github.com/alkosmas92/platform-go-challenge/internal/app/handlers"
	"github.com/alkosmas92/platform-go-challenge/internal/app/middleware"
	"github.com/alkosmas92/platform-go-challenge/internal/app/repository"
	"github.com/alkosmas92/platform-go-challenge/internal/app/services"
	"github.com/sirupsen/logrus"
)

func Run(logger *logrus.Logger, db *sql.DB) error {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService, logger)

	favoriteRepo := repository.NewFavoriteRepository(db)
	favoriteService := services.NewFavoriteService(favoriteRepo)
	favoriteHandler := handlers.NewFavoriteHandler(favoriteService, logger)

	http.HandleFunc("/register", userHandler.Register)
	http.HandleFunc("/login", userHandler.Login)

	http.Handle("/favorites", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})))

	logger.Info("Starting server on :8080")
	return http.ListenAndServe(":8080", nil)
}
