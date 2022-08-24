package di

import (
	"laramanpurego/internal/domain/controllers"
	"laramanpurego/internal/domain/repo"
	"laramanpurego/internal/presentation/http/handlers"
)

var (
	/* Controllers variable */
	userController *controllers.UserController
	userRepository *repo.UserRepository

	userHandler *handlers.UserHandler
)

func DB() {

}

/*********** Repositories ***********/

func UserRepository() *repo.UserRepository {
	if userRepository != nil {
		return userRepository
	}
	return userRepository
}

/*********** Domain ***********/

func UserDomain() *controllers.UserController {
	if userController != nil {
		return userController
	}
	userController = controllers.NewUserController(UserRepository())
	return userController
}

/*********** Handlers ***********/

func UserHandler() *handlers.UserHandler {
	if userHandler != nil {
		return userHandler
	}
	userHandler = handlers.NewUserHandler(UserDomain())
	return userHandler
}
