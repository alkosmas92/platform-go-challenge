package main

import (
	"log"

	"github.com/alkosmas92/platform-go-challenge/internal/app/database"
	"github.com/alkosmas92/platform-go-challenge/internal/app/logs"
	"github.com/alkosmas92/platform-go-challenge/internal/app/server"
)

func main() {
	logger, err := logs.Initialize()
	if err != nil {
		panic(err)
	}

	db, err := database.Initialize()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	if err := server.Run(logger, db); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
