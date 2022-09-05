package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/hamidteimouri/htutils/colog"
	"github.com/labstack/echo/v4"
	"laramanpurego/internal/domain/controllers"
	"laramanpurego/internal/domain/entity"
	"laramanpurego/internal/presentation/http/request"
	"laramanpurego/internal/presentation/http/response"
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

		return helpers.ResponseUnprocessableEntity(c, errs.Translate(translator))
	}

	/* sending data into user controller */
	token, err := u.ctrl.Login(req.Username, req.Password)
	if err != nil {
		return helpers.ResponseUnprocessableEntity(c, err.Error())
	}

	j := helpers.JwtToken{
		Token: token,
	}

	_, err = helpers.JwtTokenValidation(token)
	if err != nil {
		return err
	} else {
		colog.DoGreen("token accepted")
	}

	return helpers.ResponseOK(c, j)
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

func (u *UserHandler) GetUserByEmail(c echo.Context) error {
	email := c.Param("email")
	user, err := u.ctrl.GetUserByEmail(email)
	if err != nil {
		resp := response.Response{
			Msg: "user not found",
		}
		return helpers.ResponseNotFound(c, resp)
	}
	return helpers.ResponseOK(c, user)
}

func (u *UserHandler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user, err := u.ctrl.GetUserByID(id)
	if err != nil {
		resp := response.Response{
			Msg: "user not found",
		}
		return helpers.ResponseNotFound(c, resp)
	}
	return helpers.ResponseOK(c, user)
}

func (u *UserHandler) GetUserByMe(user *entity.User, c echo.Context) error {
	id := c.Param("id")
	user, err := u.ctrl.GetUserByID(id)
	if err != nil {
		resp := response.Response{
			Msg: "user not found",
		}
		return helpers.ResponseNotFound(c, resp)
	}
	return helpers.ResponseOK(c, user)
}
