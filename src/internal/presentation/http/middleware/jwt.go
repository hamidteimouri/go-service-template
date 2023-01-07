package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"goservicetemplate/internal/domain/entity"
	"goservicetemplate/internal/presentation/http/response"
	"goservicetemplate/pkg/helpers"
)

type userHandler func(user *entity.User, c echo.Context) error

func ValidateJwt(h userHandler) echo.HandlerFunc {

	return func(c echo.Context) error {
		bearer := c.Request().Header.Get("Authorization")
		token, ok := helpers.ExtractTokenFromAuthHeader(bearer)
		if ok == false {
			logrus.Debug("error while getting jwt token from header")
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
		if err != nil {
			return err
		}
		user := entity.User{
			Id: claims.ID,
		}

		return h(&user, c)
	}

}
