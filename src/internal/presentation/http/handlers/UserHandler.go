package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/hamidteimouri/htutils/colog"
	"github.com/labstack/echo/v4"
	"laramanpurego/src/cmd/di"
	"laramanpurego/src/internal/domain/controllers"
	request2 "laramanpurego/src/internal/presentation/http/request"
	"laramanpurego/src/pkg/helpers"
)

type UserHandler struct {
	ctrl *controllers.UserController
}

func NewUserHandler(ctrl *controllers.UserController) *UserHandler {
	return &UserHandler{ctrl: ctrl}
}

func (u *UserHandler) Login(c echo.Context) error {

	req := request2.UserLoginRequest{}
	err := c.Bind(&req)

	if err != nil {
		colog.DoRed(err.Error())
		return err
	}

	err = di.Validate().Struct(req)

	helpers.TranslateValidation(di.Validate(), di.Translator())

	if err != nil {

		errs := err.(validator.ValidationErrors)

		/*
			for _, e := range errs {
				// can translate each error one at a time.
				fmt.Println(e.Translate(trans))
			}
		*/

		helpers.ResponseUnprocessableEntity(c, errs.Translate(di.Translator()))
		return err
	}
	helpers.ResponseOK(c, "this_is_a_token")

	/* sending data into user controller */
	err = u.ctrl.Login(req.Username, req.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserHandler) Register(c echo.Context) error {
	req := request2.UserRegisterRequest{}
	err := c.Bind(&req)
	if err != nil {
		colog.DoRed(err.Error())
		return err
	}

	err = u.ctrl.Register(req.Name, req.Family, req.Username, req.Password)
	if err != nil {
		return err
	}

	return nil
}
