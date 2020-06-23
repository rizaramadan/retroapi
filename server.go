package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rizaramadan/retroapi/boards"
)

func main() {
	e := echo.New()
	g := e.Group("/api")
	boards.InitBoardHandler(g)
	e.Logger.Fatal(e.Start(":1323"))
}
