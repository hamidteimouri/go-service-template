package helpers

import (
	"github.com/labstack/echo/v4"
	"laramanpurego/internal/presentation/http/response"
	"net/http"
)

func ResponseUnprocessableEntity(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusUnprocessableEntity, data)
}

func ResponseOK(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func ResponseNotFound(c echo.Context, resp response.Response) error {
	return c.JSON(http.StatusNotFound, resp)

}
