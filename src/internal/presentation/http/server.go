package http

import (
	"github.com/hamidteimouri/gommon/htenvier"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"goservicetemplate/internal/presentation/http/routes"
)

func StartHttp() {
	e := echo.New()

	routes.Routes(e)
	address := htenvier.Env("HTTP_SERVER_ADDRESS")

	logrus.WithFields(logrus.Fields{
		"http_started_at": address,
	}).Debug("http server started")

	go func() {
		err := e.Start(address)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Panic("failed to serve http server")
		}
	}()

}
