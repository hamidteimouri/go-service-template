package di

import (
	"github.com/go-playground/validator/v10"
	"laramanpurego/internal/domain/controllers"
	"laramanpurego/internal/domain/repo"
	"laramanpurego/internal/presentation/http/handlers"
)

var (
	userRepository repo.UserRepository
	validate       validator.Validate
)

func DB() {

}


/*********** Repositories ***********/

func UserRepository() *repo.UserRepository {
	return &userRepository
}

/*********** Domain ***********/

func UserDomain() *controllers.UserController {
	userController := controllers.NewUserController(UserRepository())
	return userController
}

/*********** Handlers ***********/

func UserHandler() *handlers.UserHandler {
	uh := handlers.NewUserHandler(UserDomain())
	return uh
}
