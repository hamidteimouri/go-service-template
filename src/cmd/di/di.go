package di

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/fa"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	fa_translations "github.com/go-playground/validator/v10/translations/fa"
	"laramanpurego/internal/domain/controllers"
	"laramanpurego/internal/domain/repo"
	"laramanpurego/internal/presentation/http/handlers"
	"laramanpurego/pkg/helpers"
)

var (
	userRepository repo.UserRepository
	trans          map[string]locales.Translator
	translator     ut.Translator
	//validate       validator.Validate
)

func DB() {

}

func Translator() ut.Translator {
	fa := fa.New()
	uni := ut.New(fa, fa)
	translator, _ = uni.GetTranslator("fa")
	return translator
}
func Validate() *validator.Validate {
	validate := validator.New()
	fa_translations.RegisterDefaultTranslations(validate, translator)
	helpers.TranslateValidation(validate, translator)

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
