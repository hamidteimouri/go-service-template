package http

import (
	"github.com/labstack/echo/v4"
	"laramanpurego/internal/presentation/http/routes"
)

func StartHttp() {
	e := echo.New()

	routes.Routes(e)

	e.Start(":1323")
}
