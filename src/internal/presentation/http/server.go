package http

import (
	"fmt"
	"github.com/hamidteimouri/htutils/htcolog"
	"github.com/hamidteimouri/htutils/htenvier"
	"goservicetemplate/internal/presentation/http/routes"
)

func StartHttp() {
	e := echo.New()

	routes.Routes(e)
	addr := htenvier.Env("HTTP_SERVER_ADDRESS")
	port := htenvier.Env("HTTP_SERVER_PORT")
	address := fmt.Sprintf("%s:%s", addr, port)

	go func() {
		err := e.Start(address)
		if err != nil {
			e := fmt.Sprintf("faild to start HTTP server : %s", htcolog.MakeRed(err.Error()))
			panic(e)
		}
	}()

}
