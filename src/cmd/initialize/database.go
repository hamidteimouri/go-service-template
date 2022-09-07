package initialize

import (
	"fmt"
	"github.com/hamidteimouri/htutils/colog"
	"github.com/hamidteimouri/htutils/envier"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DatabaseInitialization() (db *gorm.DB) {
	if envier.Env("RUN_WITHOUT_DB") == "true" {
		colog.DoYellow(" - The project was executed without database check")
		return nil
	}
	var err error
	dbConnection := envier.EnvToLower("DB_CONNECTION")
	dbHost := envier.Env("DB_HOST")
	dbPort := envier.Env("DB_PORT")
	dbName := envier.Env("DB_NAME")
	dbUsername := envier.Env("DB_USERNAME")
	dbPassword := envier.Env("DB_PASSWORD")
	dbTimezone := envier.Env("DB_TIMEZONE")

	if dbConnection == "mysql" {
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			dbUsername, dbPassword, dbHost, dbPort, dbName)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	} else if dbConnection == "postgres" {
		dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v",
			dbHost, dbUsername, dbPassword, dbName, dbPort, dbTimezone)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	} else {
		panic(colog.MakeRed("invalid DB_CONNECTION (only mysql and postgres)"))
	}

	if err != nil {
		panic(colog.MakeRed("database connection error: " + err.Error()))
	} else {
		colog.DoBlue(" - successful connection to the database " + "( " + dbConnection + " )")
	}
	return

}
