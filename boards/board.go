package boards

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

//TODO: refactor this to be separated from handlers
type Board struct {
	Name        string
	Description string
}

func InitBoardApi(g *echo.Group) error {
	g.GET("/boards", List)
	return nil
}

//TODO: get this data from some sort of repository
func List(c echo.Context) error {
	var b [3]*Board
	b[0] = &Board {
		Name:  "Jon",
		Description: "this is jon board",
	}
	b[1] = &Board {
		Name:  "Doe",
		Description: "this is doe board",
	}
	b[2] = &Board {
		Name:  "Tiger",
		Description: "this is tiger board",
	}
	return c.JSONPretty(http.StatusOK, b, " ")
}
