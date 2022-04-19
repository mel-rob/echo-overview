package api

import (
	"handlers"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	e.GET("/health-check", handlers.HealthCheck)

	e.GET("/gophers/:data", handlers.GetGophers)
	e.POST("/gophers", handlers.AddGopher)

}

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
