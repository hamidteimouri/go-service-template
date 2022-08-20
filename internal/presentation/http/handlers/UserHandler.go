package handlers

import (
	"fmt"
	"github.com/hamidteimouri/htutils/colog"
	"github.com/labstack/echo/v4"
	"laramanpurego/internal/domain/controllers"
	"laramanpurego/internal/presentation/http/request"
)

type UserHandler struct {
	ctrl *controllers.UserController
}

func NewUserHandler(ctrl *controllers.UserController) *UserHandler {
	return &UserHandler{ctrl: ctrl}
}

func (u *UserHandler) Login(c echo.Context) error {
	req := request.UserLoginRequest{}
	fmt.Println(req)
	err := c.Bind(&req)
	if err != nil {
		colog.DoRed(err.Error())
		return err
	}

	err = u.ctrl.Login(req.Username, req.Password)
	if err != nil {
		return err
	}

	return nil
}
