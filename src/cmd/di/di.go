package di

import (
	"gorm.io/gorm"
	"laramanpurego/internal/data"
	"laramanpurego/internal/domain/controllers"
	"laramanpurego/internal/domain/repo"
	"laramanpurego/internal/presentation/http/handlers"
	"laramanpurego/pkg/config"
)

var (
	db *gorm.DB

	/* Controllers variable */
	userController *controllers.UserController
	userRepository repo.UserRepository

	userHandler *handlers.UserHandler
)

func DB() *gorm.DB {
	if db != nil {
		return db
	}
	db = config.DatabaseInitialization()
	return db
}

/*********** Repositories ***********/

func UserRepository() repo.UserRepository {
	if userRepository != nil {
		return userRepository
	}
	userRepository = data.NewUserRepository()
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
