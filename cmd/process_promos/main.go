package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"oolio-backend-challenge/pkg/promocode"

	_ "github.com/lib/pq"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://localhost/food_ordering?sslmode=disable"
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create context
	ctx := context.Background()

	files := []string{
		"couponbase1.gz",
		"couponbase2.gz",
		"couponbase3.gz",
	}

	for _, file := range files {
		log.Printf("Processing file: %s", file)
		if err := promocode.ProcessFile(ctx, db, file, filepath.Base(file)); err != nil {
			log.Printf("Error processing file %s: %v", file, err)
			continue
		}
		log.Printf("Finished processing file: %s", file)
	}
}
