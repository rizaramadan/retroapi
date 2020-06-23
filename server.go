package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rizaramadan/retroapi/boards"
	"net/http"
)

var (
	prefix = "/api"
	port   = ":1323"
)

func main() {
	e := echo.New()
	g := e.Group(prefix)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{ "http://localhost:8080" },
		AllowMethods: []string{ http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete },
	}))
	boards.InitBoardApi(g)
	e.Logger.Fatal(e.Start(port))
}
