package routes

import (
	"github.com/labstack/echo/v4"
	"goservicetemplate/cmd/di"
)

func Routes(e *echo.Echo) {
	/* handlers */
	userHandler := di.UserHandler()

	// initiate a route group
	g = e.Group("api")
	RegisterUserRoutes(g, userHandler)
}
