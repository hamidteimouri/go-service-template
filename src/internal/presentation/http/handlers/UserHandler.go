package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/hamidteimouri/htutils/colog"
	"github.com/labstack/echo/v4"
	"laramanpurego/internal/domain/controllers"
	"laramanpurego/internal/presentation/http/request"
	"laramanpurego/pkg/helpers"
)

type UserHandler struct {
	ctrl *controllers.UserController
}

func NewUserHandler(ctrl *controllers.UserController) *UserHandler {
	return &UserHandler{ctrl: ctrl}
}

func (u *UserHandler) Login(c echo.Context) error {

	req := request.UserLoginRequest{}
	err := c.Bind(&req)

	if err != nil {
		colog.DoRed(err.Error())
		return err
	}

	translator := helpers.Translator()
	err = helpers.Validate(translator).Struct(req)

	if err != nil {

		errs := err.(validator.ValidationErrors)

		/*
			for _, e := range errs {
				// can translate each error one at a time.
				fmt.Println(e.Translate(trans))
			}
		*/

		helpers.ResponseUnprocessableEntity(c, errs.Translate(translator))
		return nil
	}

	/* sending data into user controller */
	token, err := u.ctrl.Login(req.Username, req.Password)
	if err != nil {
		helpers.ResponseUnprocessableEntity(c, err.Error())
		return err
	}

	j := helpers.JwtToken{
		Token: token,
	}

	err = helpers.JwtTokenValidation(token)
	if err != nil {
		return err
	} else {
		colog.DoGreen("token accepted")
	}

	helpers.ResponseOK(c, j)
	return nil
}

func (u *UserHandler) Register(c echo.Context) error {
	req := request.UserRegisterRequest{}
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

func (u *UserHandler) Me(c echo.Context) (string, error) {
	return "", nil
}
