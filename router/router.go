package router

import (
	"fmt"

	"github.com/mel-rob/echo-overview/api"
	"github.com/mel-rob/echo-overview/api/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//create groups
	adminGroup := e.Group("/admin")

	//set all middlewares
	middlewares.SetMainMiddleWares(e)
	middlewares.SetAdminMiddlewares(adminGroup)

	//set main routes
	api.MainGroup(e)

	//set groupRoutes
	api.AdminGroup(adminGroup)
	fmt.Println(e.Routes())
	return e
}
