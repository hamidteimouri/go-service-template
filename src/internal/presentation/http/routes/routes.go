package routes

import (
	"github.com/labstack/echo/v4"
	"goservicetemplate/cmd/di"
	"goservicetemplate/internal/presentation/http/middleware"
)

func Routes(e *echo.Echo) {

	/*
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello,world")
		})
	*/

	/* Authentication */
	e.POST("auth/login", di.UserHandler().Login)
	e.POST("auth/register", di.UserHandler().Register)

	/* User */
	e.GET("user/email/:email", di.UserHandler().GetUserByEmail)
	e.GET("user/id/:id", di.UserHandler().GetUserByID)

	g := e.Group("/user")
	g.GET("/", middleware.ValidateJwt(di.UserHandler().Me))
	g.PATCH("/update", middleware.ValidateJwt(di.UserHandler().Update))
	g.PATCH("/password/update", middleware.ValidateJwt(di.UserHandler().UpdatePassword))

	e.GET("all", di.UserHandler().GetAll)

}
