package di

import (
	"gorm.io/gorm"
	"laramanpurego/internal/data"
	"laramanpurego/internal/data/database"
	"laramanpurego/internal/data/database/mysql"
	"laramanpurego/internal/domain/controllers"
	"laramanpurego/internal/domain/repo"
	"laramanpurego/internal/presentation/http/handlers"
	"laramanpurego/pkg/config"
)

var (
	db *gorm.DB

	/* Controllers variable */
	userController *controllers.UserController
	userRepository repo.UserRepository
	dbDatasource   database.DbDatasourceInterface

	userHandler *handlers.UserHandler
)

func DB() *gorm.DB {
	if db != nil {
		return db
	}
	db = config.DatabaseInitialization()
	return db
}

/*********** Datasource ***********/

func DbDatasource() database.DbDatasourceInterface {
	if dbDatasource != nil {
		return dbDatasource
	}
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
