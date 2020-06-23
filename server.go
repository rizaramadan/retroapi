package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	g := e.Group("/api")
	//boards.Init(g)
	e.Logger.Fatal(e.Start(":1323"))
}
