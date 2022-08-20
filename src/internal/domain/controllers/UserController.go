package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func ShowUser(c echo.Context) error {
	fmt.Println("user_id : ", c.Param("id"))

	return nil
}
