package di

import (
	"fmt"
	"github.com/hamidteimouri/htutils/colog"
	"gorm.io/gorm"
	"laramanpurego/cmd/initialize"
	"laramanpurego/internal/data"
	"laramanpurego/internal/data/database"
	"laramanpurego/internal/data/database/mysql"
	"laramanpurego/internal/domain/controllers"
	"laramanpurego/internal/domain/repo"
	"laramanpurego/internal/presentation/http/handlers"
)

var (
	db *gorm.DB

	/* Controllers variable */
	userController *controllers.UserController
	userRepository repo.UserRepository
	dbDatasource   database.DbDatasourceInterface

	/* Handlers variable */
	userHandler *handlers.UserHandler
)

func DB() *gorm.DB {
	if db != nil {
		return db
	}

	db = initialize.DatabaseInitialization()
	return db
}

/*********** Datasource ***********/

func DbDatasource() database.DbDatasourceInterface {
	if dbDatasource != nil {
		return dbDatasource
	}
	colog.DoBgYellow("in di")
	fmt.Println(db)
	dbDatasource = mysql.NewMysql(db)
	return dbDatasource
}

/*********** Repositories ***********/

func UserRepository() repo.UserRepository {
	if userRepository != nil {
		return userRepository
	}
	userRepository = data.NewUserRepository(DbDatasource())
	return userRepository
}

/*********** Domain ***********/

func UserDomain() *controllers.UserController {
	if userController != nil {
		return userController
	}
	userController = controllers.NewUserController(UserRepository())
	return userController
}

/*********** Handlers ***********/

func UserHandler() *handlers.UserHandler {
	if userHandler != nil {
		return userHandler
	}
	userHandler = handlers.NewUserHandler(UserDomain())
	return userHandler
}
