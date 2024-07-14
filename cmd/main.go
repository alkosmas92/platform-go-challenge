package main

import (
	"context"
	"github.com/alkosmas92/platform-go-challenge/internal/app/database"
	"github.com/alkosmas92/platform-go-challenge/internal/app/handlers"
	"github.com/alkosmas92/platform-go-challenge/internal/app/middleware"
	"github.com/alkosmas92/platform-go-challenge/internal/app/repository"
	"github.com/alkosmas92/platform-go-challenge/internal/app/services"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	// Set up logging
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.InfoLevel)
	// Initialize database
	db, err := database.Initialize()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	// Set up Redis client
	redisAddr := os.Getenv("REDIS_ADDR")
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr, // use environment variable
		DB:   0,         // use default DB
	})

	// Check Redis connection
	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	// Set up repositories, services, and handlers
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService, logger)

	favoriteRepo := repository.NewFavoriteRepository(db)
	favoriteService := services.NewFavoriteService(favoriteRepo)
	favoriteHandler := handlers.NewFavoriteHandler(favoriteService, logger)

	// Define routes and handlers
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

	// Start the server
	logger.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
