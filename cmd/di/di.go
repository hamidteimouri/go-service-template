package di

import (
	"laramanpurego/internal/domain/controllers"
	"laramanpurego/internal/presentation/http/handlers"
)

func DB() {

}

func UserDomain() *controllers.UserController {
	userController := controllers.NewUserController("sss")
	return userController
}

func UserHandler() *handlers.UserHandler {
	uh := handlers.NewUserHandler(UserDomain())
	return uh
}
