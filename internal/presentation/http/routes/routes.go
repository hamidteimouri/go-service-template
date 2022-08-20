package routes

import (
	"github.com/labstack/echo/v4"
	"laramanpurego/cmd/di"
	"laramanpurego/internal/domain/controllers"
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
	e.POST("auth/register", controllers.Register)

	/* User's panel */
	e.GET("user", controllers.ShowUser)
	e.POST("user", controllers.ShowUser)
	e.PATCH("user/password/change", controllers.ShowUser)
	e.GET("user/address", controllers.ShowUser)
	e.POST("user/address", controllers.ShowUser)
	e.PATCH("user/address/:id", controllers.ShowUser)

	/* Articles */
	e.POST("/article", controllers.StoreArticle)
	e.GET("/article/:slug", controllers.GetArticle)
}
