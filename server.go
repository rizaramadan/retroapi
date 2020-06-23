package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rizaramadan/retroapi/boards"
)

var (
	prefix string = "/api"
	port string = ":1323"
)

func main() {
	e := echo.New()
	g := e.Group(prefix)
	boards.InitBoardApi(g)
	e.Logger.Fatal(e.Start(port))
}
