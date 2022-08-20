package routes

import (
	"github.com/labstack/echo/v4"
	"laramanpurego/cmd/di"
	controllers2 "laramanpurego/internal/domain/controllers"
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

	/* User's panel */
	e.GET("user", controllers2.ShowUser)
	e.POST("user", controllers2.ShowUser)
	e.PATCH("user/password/change", controllers2.ShowUser)
	e.GET("user/address", controllers2.ShowUser)
	e.POST("user/address", controllers2.ShowUser)
	e.PATCH("user/address/:id", controllers2.ShowUser)

	/* Articles */
	e.POST("/article", controllers2.StoreArticle)
	e.GET("/article/:slug", controllers2.GetArticle)
}
