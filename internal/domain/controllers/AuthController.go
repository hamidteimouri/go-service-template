package controllers

import (
	"fmt"
	helpers2 "laramanpurego/pkg/helpers"

	//"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/fa"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	//en_translations "github.com/go-playground/validator/v10/translations/en"
	fa_translations "github.com/go-playground/validator/v10/translations/fa"
	"github.com/labstack/echo/v4"
)

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type UserRegister struct {
	Name     string `json:"name" validate:"required"`
	Family   string `json:"family" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserController struct {
	repo string
}

func NewUserController(repo string) *UserController {
	return &UserController{repo: repo}
}
//var tras map[string] locales.Translator = {
//	"fa" :fa.New(),
//}

func(u *UserController) Login(username, password string) error {
	//u.repo.GetUserByUsername(username)


	//en := en.New()
	//uni := ut.New(en, en)
	//trans, _ := uni.GetTranslator("en")
	//tras["fa"]
	fa := fa.New()
	uni := ut.New(fa, fa)
	trans, _ := uni.GetTranslator("fa")

	validate := validator.New()
	fa_translations.RegisterDefaultTranslations(validate, trans)

	//email := c.FormValue("email")
	//password := c.FormValue("password")

	//user := UserLogin{
	//	Email:    email,
	//	Password: password,
	//}

	//err := validate.Struct(user)

	//helpers2.TranslateValidation(validate, trans)
	//
	//if err != nil {
	//
	//	errs := err.(validator.ValidationErrors)
	//
	//	/*
	//		for _, e := range errs {
	//			// can translate each error one at a time.
	//			fmt.Println(e.Translate(trans))
	//		}
	//	*/
	//
	//	helpers2.ResponseUnprocessableEntity(c, errs.Translate(trans))
	//	return err
	//}
	//helpers2.ResponseOK(c, "this_is_a_token")

	return nil
}

func Register(c echo.Context) error {
	name := c.FormValue("name")
	family := c.FormValue("family")
	email := c.FormValue("email")
	password := c.FormValue("password")

	user := UserRegister{
		name, family, email, password,
	}

	fa := fa.New()
	uni := ut.New(fa, fa)
	trans, _ := uni.GetTranslator("fa")

	validate := validator.New()
	fa_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(user)

	helpers2.TranslateValidation(validate, trans)

	if err != nil {

		errs := err.(validator.ValidationErrors)
		helpers2.ResponseUnprocessableEntity(c, errs.Translate(trans))
		return err
	}
	helpers2.ResponseOK(c, "this_is_a_token")

	fmt.Println("email : ", email, "password : ", password, "name : ", name, "family: ", family)
	return nil
}
