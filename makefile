# Makefile for Go project

# Variables
DB_INIT_SCRIPT := internal/scripts/init_favorites.sh
MAIN_GO := cmd/main.go

# Default target
all: run

initdb:
	@echo "Running database initialization script..."
	@chmod +x $(DB_INIT_SCRIPT)
	@./$(DB_INIT_SCRIPT)

#Run the test
test:
	@go test ./...

run:
	@echo "Running the Go application..."
	@go run $(MAIN_GO)

# Clean up the database file
clean:
	@echo "Cleaning up..."
	@rm -f scripts/favorites.db

