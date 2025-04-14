package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lib/pq"

	"photogo/internal/database"
	"photogo/internal/templates/pages"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Initialize database
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Construct database URL from individual parameters
		dbURL = "postgres://" + getEnv("DB_USER", "postgres") + ":" +
			getEnv("DB_PASSWORD", "postgres") + "@" +
			getEnv("DB_HOST", "localhost") + ":" +
			getEnv("DB_PORT", "5432") + "/" +
			getEnv("DB_NAME", "photogo") + "?sslmode=" +
			getEnv("DB_SSLMODE", "disable")
	}

	// Parse the database URL
	pgURL, err := pq.ParseURL(dbURL)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
	}

	// Initialize database connection
	db, err := database.New(pgURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run migrations
	if err := db.Migrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static files
	e.Static("/static", "static")

	// Routes
	e.GET("/", func(c echo.Context) error {
		return pages.Home().Render(context.Background(), c.Response().Writer)
	})

	e.POST("/api/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from HTMX!")
	})

	// Start server
	port := getEnv("PORT", "6969")
	e.Logger.Fatal(e.Start(":" + port))
} 