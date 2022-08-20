package database

import (
	"fmt"
	"github.com/hamidteimouri/htutils/colog"
	"github.com/hamidteimouri/htutils/envier"
)

func Start() {
	connection()
}
func connection() {
	dbConnection := envier.EnvToLower("DB_CONNECTION")
	dbHost := envier.EnvToLower("DB_HOST")
	dbPort := envier.EnvToLower("DB_PORT")
	dbName := envier.EnvToLower("DB_NAME")
	dbUsername := envier.EnvToLower("DB_USERNAME")
	dbPassword := envier.Env("DB_PASSWORD")

	if envier.EnvExists("DB_CONNECTION") == false {
		colog.DoRed("DB_CONNECTION is required")
	}
	if envier.EnvExists("DB_HOST") == false {
		colog.DoRed("DB_HOST is required")
	}
	if envier.EnvExists("DB_PORT") == false {
		colog.DoRed("DB_PORT is required")
	}
	if envier.EnvExists("DB_NAME") == false {
		colog.DoRed("DB_NAME is required")
	}
	if envier.EnvExists("DB_USERNAME") == false {
		colog.DoRed("DB_USERNAME is required")
	}
	if envier.EnvExists("DB_PASSWORD") == false {
		colog.DoRed("DB_PASSWORD is required")
	}

	fmt.Println("DB Connection: ", dbConnection)
	fmt.Println("DB Host: ", dbHost)
	fmt.Println("DB Port: ", dbPort)
	fmt.Println("DB Name: ", dbName)
	fmt.Println("DB Username: ", dbUsername)
	fmt.Println("DB Password: ", dbPassword)
}
