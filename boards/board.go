package boards

import (
	"github.com/labstack/echo"
	"net/http"
)

type Boards struct { }

func InitBoardHandler(g echo.Group) error {
	g.GET("boards", List)
}

func List(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}