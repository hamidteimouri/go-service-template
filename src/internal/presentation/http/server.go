package http

import (
	"github.com/hamidteimouri/htutils/colog"
	"github.com/labstack/echo/v4"
	"laramanpurego/internal/presentation/http/routes"
)

func StartHttp() {
	e := echo.New()

	routes.Routes(e)

	err := e.Start(":1323")
	if err != nil {
		colog.DoBgRed(err.Error())
		return
	}
}
