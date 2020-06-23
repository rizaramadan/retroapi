package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rizaramadan/retroapi/boards"
	"github.com/rizaramadan/retroapi/internal"
)

var (
	prefix = "/api"
	port   = ":1323"
)

func main() {
	e := echo.New()
	internal.InitEcho(e)
	g := e.Group(prefix)

	err := boards.InitBoardApi(g)
	if err != nil {
		fmt.Println(fmt.Errorf("error InitBoardApi: %v", err))
	}

	e.Logger.Fatal(e.Start(port))
}
