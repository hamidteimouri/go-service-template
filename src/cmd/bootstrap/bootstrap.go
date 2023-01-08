package bootstrap

import (
	"github.com/hamidteimouri/gommon/htcolog"
	"github.com/hamidteimouri/gommon/htenvier"
	"github.com/sirupsen/logrus"
	"goservicetemplate/cmd/di"
	model "goservicetemplate/internal/data/database/postgres"
	"goservicetemplate/internal/presentation/grpc"
	"goservicetemplate/internal/presentation/http"
)

func Start() {
	checkImportantEnv()
	databaseMigration()
	http.StartHttp()
	grpc.StartGRPC()
}

func databaseMigration() {
	if htenvier.Env("RUN_WITHOUT_DB") == "true" {
		return
	}
	err := di.DB().AutoMigrate(model.UserModel{})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("failed to run migrations")
	}

}

func checkImportantEnv() {

	needToPanic := false
	if htenvier.EnvExists("DB_CONNECTION") == false {
		logrus.Error("DB_CONNECTION is required")
		needToPanic = true
	}
	if htenvier.EnvExists("DB_HOST") == false {
		logrus.Error("DB_HOST is required")
		needToPanic = true
	}
	if htenvier.EnvExists("DB_PORT") == false {
		logrus.Error("DB_PORT is required")
		needToPanic = true
	}
	if htenvier.EnvExists("DB_NAME") == false {
		logrus.Error("DB_NAME is required")
		needToPanic = true
	}
	if htenvier.EnvExists("DB_USERNAME") == false {
		logrus.Error("DB_USERNAME is required")
		needToPanic = true
	}
	if htenvier.EnvExists("DB_PASSWORD") == false {
		logrus.Error("DB_PASSWORD is required")
		needToPanic = true
	}
	if htenvier.EnvExists("JWT_SIGNING_KEY") == false {
		logrus.Error("JWT_SIGNING_KEY is required")
		needToPanic = true
	}
	if htenvier.EnvExists("JWT_EXPIRE_MINUTES") == false {
		logrus.Error("JWT_EXPIRE_MINUTES is required")
		needToPanic = true
	}
	if htenvier.EnvExists("TIMEZONE") == false {
		logrus.Error("TIMEZONE is required")
		needToPanic = true
	}
	if htenvier.EnvExists("APP_ENV") == false {
		logrus.Error("APP_ENV is required")
		needToPanic = true
	}
	if htenvier.EnvExists("HTTP_SERVER_ADDRESS") == false {
		logrus.Error("HTTP_SERVER_ADDRESS is required")
		needToPanic = true
	}

	if htenvier.EnvExists("GRPC_SERVER_ADDRESS") == false {
		logrus.Error("GRPC_SERVER_ADDRESS is required")
		needToPanic = true
	}

	if needToPanic {
		panic(htcolog.MakeRed("complete the ENV file"))
	}

}
