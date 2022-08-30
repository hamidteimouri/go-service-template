package routes

import (
	"github.com/labstack/echo/v4"
	"laramanpurego/cmd/di"
	"net/http"
)

func Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello,world")
	})

	//userController := controllers.NewUserController("sss")
	//userHandler := handlers.NewUserHandler(userController)

	/* Authentication */
	e.POST("auth/login", di.UserHandler().Login)
	e.POST("auth/register", di.UserHandler().Register)

	//g := e.Group("/user")
	//g.GET("/", middleware.ValidateJwt(di.UserHandler()))

}
