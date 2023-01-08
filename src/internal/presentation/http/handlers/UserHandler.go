package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/hamidteimouri/gommon/htcolog"
	"github.com/labstack/echo/v4"
	"goservicetemplate/internal/domain/controllers"
	"goservicetemplate/internal/domain/entity"
	"goservicetemplate/internal/presentation/http/dto"
	"goservicetemplate/internal/presentation/http/request"
	"goservicetemplate/internal/presentation/http/response"
	"goservicetemplate/pkg/helpers"
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
		htcolog.DoRed(err.Error())
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
		resp := response.Response{
			Data: errs.Translate(translator),
		}
		return helpers.ResponseUnprocessableEntity(c, resp)
	}

	/* sending data into user controller */
	token, err := u.ctrl.Login(req.Username, req.Password)
	if err != nil {
		resp := response.Response{
			Msg: err.Error(),
		}
		return helpers.ResponseUnprocessableEntity(c, resp)
	}
	j := helpers.JwtToken{
		Token: token,
	}
	resp := response.Response{
		Data: j,
	}

	return helpers.ResponseOK(c, resp)
	/*
		j := helpers.JwtToken{
			Token: token,
		}

		_, err = helpers.JwtTokenValidation(token)
		if err != nil {
			colog.DoBgRed(err.Error())
			return err
		} else {
			colog.DoGreen("token accepted")
		}
	*/

	//return helpers.ResponseOK(c, j)
}

func (u *UserHandler) Register(c echo.Context) error {
	req := request.UserRegisterRequest{}
	err := c.Bind(&req)
	if err != nil {
		return err
	}
	translator := helpers.Translator()
	err = helpers.Validate(translator).Struct(req)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		resp := response.Response{
			Data: errs.Translate(translator),
		}
		return helpers.ResponseUnprocessableEntity(c, resp)
	}

	err = u.ctrl.Register(req.Name, req.Family, req.Username, req.Password)
	if err != nil {
		resp := response.Response{
			Msg: "error: " + err.Error(),
		}
		return helpers.ResponseUnprocessableEntity(c, resp)
	}
	resp := response.Response{
		Msg: "ثبت نام شما با موفقیت انجام شد",
	}

	return helpers.ResponseOK(c, resp)
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
	resp := response.Response{
		Msg:  "",
		Data: user,
	}
	return helpers.ResponseOK(c, resp)
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
	resp := response.Response{
		Data: user,
	}
	return helpers.ResponseOK(c, resp)
}

func (u *UserHandler) Me(user *entity.User, c echo.Context) error {
	user, err := u.ctrl.GetUserByID(user.Id)
	if err != nil {
		resp := response.Response{
			Msg: "user not found",
		}
		return helpers.ResponseNotFound(c, resp)
	}
	userDto := dto.User{}
	userDto.ConvertEntityToDTO(user)
	resp := response.Response{
		Data: userDto,
	}
	return helpers.ResponseOK(c, resp)
}

func (u *UserHandler) UpdatePassword(user *entity.User, c echo.Context) error {
	user, err := u.ctrl.GetUserByID(user.Id)
	if err != nil {
		resp := response.Response{
			Msg: "user not found",
		}
		return helpers.ResponseNotFound(c, resp)
	}

	req := request.UserChangePasswordRequest{}
	err = c.Bind(&req)
	if err != nil {
		resp := response.Response{
			Msg: "internal error",
		}
		return helpers.ResponseInternalError(c, resp)
	}
	err = nil
	translator := helpers.Translator()
	err = helpers.Validate(translator).Struct(req)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		resp := response.Response{
			Data: errs.Translate(translator),
		}
		return helpers.ResponseUnprocessableEntity(c, resp)
	}
	ok := helpers.HashCheck(req.CurrentPassword, user.Password)
	if ok == false {
		resp := response.Response{
			Msg: "password is wrong",
		}
		return helpers.ResponseUnprocessableEntity(c, resp)
	}

	ok, err = u.ctrl.ChangePassword(user, req.Password)
	if err != nil || ok == false {
		resp := response.Response{
			Data: err.Error(),
		}
		return helpers.ResponseUnprocessableEntity(c, resp)
	}

	userDto := dto.User{}
	userDto.ConvertEntityToDTO(user)
	resp := response.Response{
		Data: userDto,
	}
	return helpers.ResponseOK(c, resp)
}

func (u *UserHandler) Update(user *entity.User, c echo.Context) error {
	user, err := u.ctrl.GetUserByID(user.Id)
	if err != nil {
		resp := response.Response{
			Msg: "user not found",
		}
		return helpers.ResponseNotFound(c, resp)
	}

	req := request.UserUpdateRequest{}
	err = c.Bind(&req)
	if err != nil {
		resp := response.Response{
			Msg: "internal error",
		}
		return helpers.ResponseInternalError(c, resp)
	}
	err = nil
	translator := helpers.Translator()
	err = helpers.Validate(translator).Struct(req)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		resp := response.Response{
			Data: errs.Translate(translator),
		}
		return helpers.ResponseUnprocessableEntity(c, resp)
	}

	user.Name = req.Name
	user.Family = req.Family
	user.Mobile = req.Mobile

	result, ok, err := u.ctrl.Update(user)
	if err != nil || ok == false {
		resp := response.Response{
			Data: err.Error(),
		}
		return helpers.ResponseUnprocessableEntity(c, resp)
	}

	userDto := dto.User{}
	userDto.ConvertEntityToDTO(result)
	resp := response.Response{
		Data: userDto,
	}
	return helpers.ResponseOK(c, resp)
}

func (u *UserHandler) GetAll(c echo.Context) error {

	ch := u.ctrl.GetAll()

	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().Header().Set("Transfer-Encoding", "chunked")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

Loop:
	for {
		select {
		case result, ok := <-ch:
			if !ok {
				fmt.Println("closed")
				break Loop
			}
			if result.Error != nil {
				fmt.Println(result.Error)
				break Loop
			}

			fmt.Println(result)
			user, err := json.Marshal(result)
			if err != nil {
				break Loop
			}
			_, err = fmt.Fprintf(c.Response().Writer, "data:%s\n\n", string(user))
			if err != nil {

				break Loop
			}
			c.Response().Flush()

		case <-c.Request().Context().Done():
			break Loop
		}
	}
	return nil

}
