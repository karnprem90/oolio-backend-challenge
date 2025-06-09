package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"oolio-backend-challenge/internal/handler"
	"oolio-backend-challenge/internal/repository/postgres"
	"oolio-backend-challenge/internal/router"
	"oolio-backend-challenge/internal/service"
)

func main() {
	// Initialize database connection
	db, err := sqlx.Connect("postgres", "dbname=food_ordering sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	productRepo := postgres.NewProductRepository(db)
	orderRepo := postgres.NewOrderRepository(db)
	promoRepo := postgres.NewPromoRepository(db)

	// Initialize services
	productService := service.NewProductService(productRepo)
	orderService := service.NewOrderService(orderRepo, productRepo)
	promoService := service.NewPromoCodeService(promoRepo)

	// Initialize handlers
	productHandler := handler.NewProductHandler(productService)
	orderHandler := handler.NewOrderHandler(orderService)
	promoHandler := handler.NewPromoHandler(promoService)

	// Initialize and setup router
	r := router.New()
	r.RegisterProductRoutes(productHandler)
	r.RegisterOrderRoutes(orderHandler)
	r.RegisterPromoRoutes(promoHandler)

	// Get port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	// Create server with timeouts
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r.Engine(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}

func getDBConnString() string {
	return getEnv("DATABASE_URL", "postgres://localhost:5432/food_ordering?sslmode=disable")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
