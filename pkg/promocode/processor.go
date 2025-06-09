package promocode

import (
	"bufio"
	"compress/gzip"
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

func ProcessFile(ctx context.Context, db *sql.DB, filePath string, fileName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzReader.Close()

	reader := bufio.NewReader(gzReader)

	stmt, err := db.PrepareContext(ctx, `
		INSERT INTO promo_codes (code, file_name)
		VALUES ($1, $2)
		ON CONFLICT (code, file_name) DO NOTHING
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	const batchSize = 1000
	var codes []string
	var mu sync.Mutex

	validCodes := make(chan string, batchSize)

	go func() {
		for code := range validCodes {
			mu.Lock()
			codes = append(codes, code)
			if len(codes) >= batchSize {
				// Insert batch
				for _, c := range codes {
					_, err := tx.Stmt(stmt).ExecContext(ctx, c, fileName)
					if err != nil {
						log.Printf("Error inserting code %s: %v\n", c, err)
					} else {
						log.Printf("Inserted promo code: %s from file: %s\n", c, fileName)
					}
				}
				codes = codes[:0]
			}
			mu.Unlock()
		}
	}()

	lineCount := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error reading file: %w", err)
		}

		code := strings.TrimSpace(line)
		if isValidCode(code) {
			validCodes <- code
		}

		lineCount++
		if lineCount%10000 == 0 {
			log.Printf("Processed %d lines from %s\n", lineCount, fileName)
		}

		if lineCount >= 5000 {
			log.Printf("Reached 5000 lines limit for file: %s\n", fileName)
			break
		}
	}

	close(validCodes)

	mu.Lock()
	if len(codes) > 0 {
		for _, c := range codes {
			_, err := tx.Stmt(stmt).ExecContext(ctx, c, fileName)
			if err != nil {
				log.Printf("Error inserting code %s: %v\n", c, err)
			} else {
				log.Printf("Inserted promo code: %s from file: %s\n", c, fileName)
			}
		}
	}
	mu.Unlock()

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func isValidCode(code string) bool {
	length := len(code)
	return length >= 8 && length <= 10
}
