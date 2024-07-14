# Variables
MAIN=cmd/main.go
TEST_DIR=./...
SCRIPT_DIR=scripts
POPULATE_DB_SCRIPT=$(SCRIPT_DIR)/populate_database.sh

# Targets
all: test run populate_db

test:
	@echo "Running tests..."
	@go test -v $(TEST_DIR)

run:
	@echo "Running the main application..."
	@go run $(MAIN) &

populate_db:
	@echo "Populating the database..."
	@sleep 1
	@bash $(POPULATE_DB_SCRIPT)

.PHONY: all test run populate_db
