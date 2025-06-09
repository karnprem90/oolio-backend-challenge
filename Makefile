.PHONY: build run test migrate-up migrate-down seed-data unseed-data process-promos

# Build the application
build:
	go build -o bin/api cmd/api/main.go
	go build -o bin/process_promos cmd/process_promos/main.go

# Run the application
run:
	go run cmd/api/main.go

# Run tests
test:
	go test -v ./...

# Run database migrations
migrate-up:
	psql -d food_ordering -f migrations/000001_init_schema.up.sql

# Revert database migrations
migrate-down:
	psql -d food_ordering -f migrations/000001_init_schema.down.sql

# Create database
create-db:
	createdb food_ordering

# Drop database
drop-db:
	dropdb food_ordering

# Seed sample data
seed-data:
	psql -d food_ordering -f migrations/000002_seed_data.up.sql

# Remove sample data
unseed-data:
	psql -d food_ordering -f migrations/000002_seed_data.down.sql

# Install dependencies
deps:
	go mod download

# Clean build artifacts
clean:
	rm -rf bin/

# Process promo code files
process-promos:
	./bin/process_promos 