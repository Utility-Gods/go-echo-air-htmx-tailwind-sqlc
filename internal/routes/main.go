package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Public Page Routes
	e.GET("/", handlers.HomeHandler)
	e.GET("/about", handlers.AboutHandler)
	e.GET("/contact", handlers.ContactHandler)

	// Protected App Routes
	appGroup := e.Group("/app")
	{
		// Dashboard Routes
		appGroup.GET("/dashboard", handlers.DashboardHandler)
		
		// Profile Routes
		appGroup.GET("/profile", handlers.ProfileHandler)
	}

	// API Routes
	apiGroup := e.Group("/api")
	{
		// Example API endpoints
		apiGroup.POST("/hello", api.HelloHandler)
		apiGroup.GET("/status", api.StatusHandler)
	}

	// Static Files
	e.Static("/static", "static")

	// Global Catch-all Route
	e.GET("/*", handlers.NotFoundHandler)
} 