package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var (
	url = "http://localhost:8080"
)

func InitEcho(e *echo.Echo)  {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig {
		AllowOrigins: []string{ url },
		AllowMethods: []string{ http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete },
	}))
}