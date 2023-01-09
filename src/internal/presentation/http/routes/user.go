package routes

import (
	"github.com/labstack/echo/v4"
	"goservicetemplate/internal/presentation/http/handlers"
	"goservicetemplate/internal/presentation/http/middleware"
)

func RegisterUserRoutes(group *echo.Group, handler *handlers.UserHandler) {
	// Authentication
	group.POST("auth/login", handler.Login)
	group.POST("auth/login/", handler.Login)
	group.POST("auth/register", handler.Register)
	group.POST("auth/register/", handler.Register)

	// User
	group.GET("/user", middleware.ValidateJwt(handler.Me))
	group.GET("/user/", middleware.ValidateJwt(handler.Me))
	group.GET("/user/", middleware.ValidateJwt(handler.Me))
	group.PATCH("/update", middleware.ValidateJwt(handler.Update))
	group.PATCH("/password/update", middleware.ValidateJwt(handler.UpdatePassword))

	// Admin Routes
	group.GET("/users", handler.GetAll)
	group.GET("user/email/:email", handler.GetUserByEmail)
	group.GET("user/id/:id", handler.GetUserByID)

}
