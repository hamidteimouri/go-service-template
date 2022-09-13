package initialize

import (
	"fmt"
	"github.com/hamidteimouri/htutils/htcolog"
	"github.com/hamidteimouri/htutils/htenvier"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DatabaseInitialization() (db *gorm.DB) {
	if htenvier.Env("RUN_WITHOUT_DB") == "true" {
		htcolog.DoYellow(" - The project was executed without database check")
		return nil
	}
	var err error
	dbConnection := htenvier.EnvToLower("DB_CONNECTION")
	dbHost := htenvier.Env("DB_HOST")
	dbPort := htenvier.Env("DB_PORT")
	dbName := htenvier.Env("DB_NAME")
	dbUsername := htenvier.Env("DB_USERNAME")
	dbPassword := htenvier.Env("DB_PASSWORD")
	dbTimezone := htenvier.Env("DB_TIMEZONE")

	if dbConnection == "mysql" {
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			dbUsername, dbPassword, dbHost, dbPort, dbName)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	} else if dbConnection == "postgres" {
		dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v",
			dbHost, dbUsername, dbPassword, dbName, dbPort, dbTimezone)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	} else {
		panic(htcolog.MakeRed("invalid DB_CONNECTION (only mysql and postgres)"))
	}

	if err != nil {
		panic(htcolog.MakeRed("database connection error: " + err.Error()))
	} else {
		htcolog.DoBlue(" - successful connection to the database " + "( " + dbConnection + " )")
	}
	return

}
