package helpers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ResponseUnprocessableEntity(c echo.Context, data interface{}) {
	c.JSON(http.StatusUnprocessableEntity, data)
}

func ResponseOK(c echo.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
