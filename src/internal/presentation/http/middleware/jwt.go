package middleware

import (
	"errors"
	"fmt"
	"github.com/hamidteimouri/htutils/colog"
	"github.com/labstack/echo/v4"
	"laramanpurego/internal/domain/entity"
	"laramanpurego/pkg/helpers"
)

type userHandler func(user *entity.User, c echo.Context) error

func ValidateJwt(h userHandler) echo.HandlerFunc {

	return func(c echo.Context) error {
		bearer := c.Request().Header.Get("Authorization")
		token, ok := helpers.ExtractTokenFromAuthHeader(bearer)
		if ok != false {
			colog.DoRed("error while getting jwt token from header")
			return errors.New("something went wrong")
		}
		claims, err := helpers.JwtTokenValidation(token)
		if err != nil {
			return err
		}
		fmt.Println("claims: ", claims)

		return h(nil, c)
	}

}
