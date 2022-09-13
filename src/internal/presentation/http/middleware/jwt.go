package middleware

import (
	"errors"
	"fmt"
	"github.com/hamidteimouri/htutils/htcolog"
	"github.com/labstack/echo/v4"
	"laramanpurego/internal/domain/entity"
	"laramanpurego/internal/presentation/http/response"
	"laramanpurego/pkg/helpers"
	"strconv"
)

type userHandler func(user *entity.User, c echo.Context) error

func ValidateJwt(h userHandler) echo.HandlerFunc {

	return func(c echo.Context) error {
		bearer := c.Request().Header.Get("Authorization")
		token, ok := helpers.ExtractTokenFromAuthHeader(bearer)
		if ok == false {
			htcolog.DoRed("error while getting jwt token from header")
			return errors.New("something went wrong")
		}
		claims, err := helpers.JwtTokenValidation(token)
		if err != nil {
			resp := response.Response{
				Msg: "unauthorized",
			}
			return helpers.ResponseUnauthorized(c, resp)
		}
		fmt.Println("claims: ", claims)
		userId, err := strconv.ParseUint(claims.ID, 10, 64)
		if err != nil {
			return err
		}
		user := entity.User{
			Id: uint(userId),
		}

		return h(&user, c)
	}

}
