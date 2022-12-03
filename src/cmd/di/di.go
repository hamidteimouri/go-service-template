package di

import (
	"gorm.io/gorm"
	"goservicetemplate/cmd/initialize"
	"goservicetemplate/internal/data"
	"goservicetemplate/internal/data/database"
	"goservicetemplate/internal/data/database/mysql"
	"goservicetemplate/internal/domain/controllers"
	"goservicetemplate/internal/domain/repo"
	"goservicetemplate/internal/presentation/grpc/servers"
	"goservicetemplate/internal/presentation/http/handlers"
)

var (
	db *gorm.DB

	/* Controllers variable */
	userController *controllers.UserController
	userRepository repo.UserRepository
	dbDatasource   database.DbDatasourceInterface

	/* Handlers variable */
	userHandler *handlers.UserHandler

	/* GRPC variable */
	userServer *servers.UserServer
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

/*********** GRPC ***********/

func GrpcUserServer() *servers.UserServer {
	if userServer != nil {
		return userServer
	}
	userServer = servers.NewUserServer(userController)
	return userServer
}
