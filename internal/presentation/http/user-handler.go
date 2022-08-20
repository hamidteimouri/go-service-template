package http

import (
	"github.com/labstack/echo/v4"
	"laramanpurego/internal/domain/controllers"
	"net/http"
)

type userHandler struct {
	ctrl *controllers.UserController
}

func newUserHandler(ctrl *controllers.UserController) *userHandler {
	return &userHandler{ctrl: ctrl}
}
func (u userHandler) login(c echo.Context) error {
	type Req struct {
		Username string `json:"username1"`
		Password string `json:"password"`
	}
	request := &Req{}
	err := c.Bind(request)
	if err != nil {
		return err
	}
	err = u.ctrl.Login(request.Username, request.Password)
	if err != nil {
		return c.JSON(http.StatusNotFound, &Response{Error: err})
	}
	//c.JSON(http.StatusOK, &Response{nil, nil})
	return nil
}

type Response struct {
	Error error       `json:"error"`
	Data  interface{} `json:"data"`
}
