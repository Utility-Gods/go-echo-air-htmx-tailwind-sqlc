package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.HTML(http.StatusOK, "<span class='font-medium'>Hello from HTMX!</span> This message was loaded dynamically.")
} 