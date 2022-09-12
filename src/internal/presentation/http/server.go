package http

import (
	"fmt"
	"github.com/hamidteimouri/htutils/colog"
	"github.com/hamidteimouri/htutils/envier"
	"github.com/labstack/echo/v4"
	"laramanpurego/internal/presentation/http/routes"
)

func StartHttp() {
	e := echo.New()

	routes.Routes(e)
	addr := envier.Env("HTTP_SERVER_ADDRESS")
	port := envier.Env("HTTP_SERVER_PORT")
	address := fmt.Sprintf("%s:%s", addr, port)

	go func() {
		err := e.Start(address)
		if err != nil {
			e := fmt.Sprintf("faild to start HTTP server : %s", colog.MakeRed(err.Error()))
			panic(e)
		}
	}()

}
