package api

import (
	"github.com/mel-rob/echo-overview/handlers"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	e.POST("/gophers", handlers.AddGopher)
	e.GET("/gophers/:data", handlers.GetGophers)

}

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
