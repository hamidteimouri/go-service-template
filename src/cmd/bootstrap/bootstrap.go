package bootstrap

import (
	"github.com/hamidteimouri/htutils/htcolog"
	"github.com/hamidteimouri/htutils/htenvier"
	"laramanpurego/cmd/di"
	model "laramanpurego/internal/data/database/mysql"
	"laramanpurego/internal/presentation/grpc"
	"laramanpurego/internal/presentation/http"
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
		return
	}

}

func checkImportantEnv() {

	needToPanic := false
	if htenvier.EnvExists("DB_CONNECTION") == false {
		htcolog.DoRed("DB_CONNECTION is required")
		needToPanic = true
	}
	if htenvier.EnvExists("DB_HOST") == false {
		htcolog.DoRed("DB_HOST is required")
		needToPanic = true
	}
	if htenvier.EnvExists("DB_PORT") == false {
		htcolog.DoRed("DB_PORT is required")
		needToPanic = true
	}
	if htenvier.EnvExists("DB_NAME") == false {
		htcolog.DoRed("DB_NAME is required")
		needToPanic = true
	}
	if htenvier.EnvExists("DB_USERNAME") == false {
		htcolog.DoRed("DB_USERNAME is required")
		needToPanic = true
	}
	if htenvier.EnvExists("DB_PASSWORD") == false {
		htcolog.DoRed("DB_PASSWORD is required")
		needToPanic = true
	}
	if htenvier.EnvExists("JWT_SIGNING_KEY") == false {
		htcolog.DoRed("JWT_SIGNING_KEY is required")
		needToPanic = true
	}
	if htenvier.EnvExists("JWT_EXPIRE_MINUTES") == false {
		htcolog.DoRed("JWT_EXPIRE_MINUTES is required")
		needToPanic = true
	}
	if htenvier.EnvExists("TIMEZONE") == false {
		htcolog.DoRed("TIMEZONE is required")
		needToPanic = true
	}
	if htenvier.EnvExists("APP_ENV") == false {
		htcolog.DoRed("APP_ENV is required")
		needToPanic = true
	}
	if htenvier.EnvExists("HTTP_SERVER_ADDRESS") == false {
		htcolog.DoRed("HTTP_SERVER_ADDRESS is required")
		needToPanic = true
	}
	if htenvier.EnvExists("HTTP_SERVER_PORT") == false {
		htcolog.DoRed("HTPP_SERVER_PORT is required")
		needToPanic = true
	}

	if htenvier.EnvExists("GRPC_SERVER_ADDRESS") == false {
		htcolog.DoRed("GRPC_SERVER_ADDRESS is required")
		needToPanic = true
	}
	if htenvier.EnvExists("GRPC_SERVER_PORT") == false {
		htcolog.DoRed("GRPC_SERVER_PORT is required")
		needToPanic = true
	}
	if needToPanic {
		panic(htcolog.MakeRed("complete the ENV file"))
	}

}
