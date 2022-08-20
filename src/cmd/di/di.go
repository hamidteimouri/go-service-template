package di

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/fa"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	fa_translations "github.com/go-playground/validator/v10/translations/fa"
	"laramanpurego/src/internal/domain/controllers"
	"laramanpurego/src/internal/domain/repo"
	"laramanpurego/src/internal/presentation/http/handlers"
)

var (
	userRepository repo.UserRepository
	trans          map[string]locales.Translator
	//validate       validator.Validate
)

//var Translation map[string]interface{} = {
//"fa":fa.New()
//}

func DB() {

}

func Translator() ut.Translator {
	fa := fa.New()
	uni := ut.New(fa, fa)
	trans, _ := uni.GetTranslator("fa")
	return trans
}
func Validate() *validator.Validate {
	validate := validator.New()
	fa_translations.RegisterDefaultTranslations(validate, Translator())

	return validate
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
