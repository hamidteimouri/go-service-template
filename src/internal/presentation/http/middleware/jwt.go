package middleware

import (
	"github.com/hamidteimouri/htutils/htcolog"
	"goservicetemplate/internal/domain/entity"
	"goservicetemplate/internal/presentation/http/response"
	"goservicetemplate/pkg/helpers"
	"strconv"
)

type userHandler func(user *entity.User, c echo.Context) error

func ValidateJwt(h userHandler) echo.HandlerFunc {

	return func(c echo.Context) error {
		bearer := c.Request().Header.Get("Authorization")
		token, ok := helpers.ExtractTokenFromAuthHeader(bearer)
		if ok == false {
			htcolog.DoRed("error while getting jwt token from header")
			resp := response.Response{
				Msg: "unauthorized",
			}
			return helpers.ResponseUnauthorized(c, resp)
		}
		claims, err := helpers.JwtTokenValidation(token)
		if err != nil {
			resp := response.Response{
				Msg: "unauthorized",
			}
			return helpers.ResponseUnauthorized(c, resp)
		}
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
