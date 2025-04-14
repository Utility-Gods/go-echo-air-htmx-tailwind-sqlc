package main

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"photogo/internal/api"
	"photogo/internal/templates/pages"
)

func main() {
	// Create a new echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve static files
	e.Static("/static", "static")

	// Pages
	e.GET("/", func(c echo.Context) error {
		return pages.Home().Render(context.Background(), c.Response().Writer)
	})

	// API routes
	apiGroup := e.Group("/api")
	apiGroup.POST("/hello", api.Hello)

	// Start server
	if err := e.Start(":6969"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
} 