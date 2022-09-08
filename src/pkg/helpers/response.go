package helpers

import (
	"github.com/labstack/echo/v4"
	"laramanpurego/internal/presentation/http/response"
	"net/http"
)

func ResponseUnprocessableEntity(c echo.Context, resp response.Response) error {
	return c.JSON(http.StatusUnprocessableEntity, resp)
}

func ResponseOK(c echo.Context, resp response.Response) error {
	return c.JSON(http.StatusOK, resp)
}

func ResponseNotFound(c echo.Context, resp response.Response) error {
	return c.JSON(http.StatusNotFound, resp)
}

func ResponseUnauthorized(c echo.Context, resp response.Response) error {
	return c.JSON(http.StatusUnauthorized, resp)
}

func ResponseInternalError(c echo.Context, resp response.Response) error {
	return c.JSON(http.StatusInternalServerError, resp)
}
