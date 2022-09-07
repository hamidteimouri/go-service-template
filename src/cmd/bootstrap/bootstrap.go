package bootstrap

import (
	"github.com/hamidteimouri/htutils/colog"
	"github.com/hamidteimouri/htutils/envier"
	"laramanpurego/cmd/di"
	model "laramanpurego/internal/data/database/mysql"
	"laramanpurego/internal/presentation/http"
)

func Start() {
	checkImportantEnv()
	databaseMigration()
	http.StartHttp()
}

func databaseMigration() {
	if envier.Env("RUN_WITHOUT_DB") == "true" {
		return
	}
	err := di.DB().AutoMigrate(model.UserModel{})
	if err != nil {
		return
	}

}

func checkImportantEnv() {

	needToPanic := false
	if envier.EnvExists("DB_CONNECTION") == false {
		colog.DoRed("DB_CONNECTION is required")
		needToPanic = true
	}
	if envier.EnvExists("DB_HOST") == false {
		colog.DoRed("DB_HOST is required")
		needToPanic = true
	}
	if envier.EnvExists("DB_PORT") == false {
		colog.DoRed("DB_PORT is required")
		needToPanic = true
	}
	if envier.EnvExists("DB_NAME") == false {
		colog.DoRed("DB_NAME is required")
		needToPanic = true
	}
	if envier.EnvExists("DB_USERNAME") == false {
		colog.DoRed("DB_USERNAME is required")
		needToPanic = true
	}
	if envier.EnvExists("DB_PASSWORD") == false {
		colog.DoRed("DB_PASSWORD is required")
		needToPanic = true
	}
	if envier.EnvExists("JWT_SIGNING_KEY") == false {
		colog.DoRed("JWT_SIGNING_KEY is required")
		needToPanic = true
	}
	if envier.EnvExists("JWT_EXPIRE_MINUTES") == false {
		colog.DoRed("JWT_EXPIRE_MINUTES is required")
		needToPanic = true
	}
	if envier.EnvExists("TIMEZONE") == false {
		colog.DoRed("TIMEZONE is required")
		needToPanic = true
	}
	if envier.EnvExists("APP_ENV") == false {
		colog.DoRed("APP_ENV is required")
		needToPanic = true
	}
	if needToPanic {
		panic(colog.MakeRed("complete the ENV file"))
	}

}
