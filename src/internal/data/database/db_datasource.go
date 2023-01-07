package database

import (
	"goservicetemplate/internal/domain/dto"
	"goservicetemplate/internal/domain/entity"
)

type DbDatasourceInterface interface {
	FindUserById(id string) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	InsertUser(user *entity.User) (*entity.User, error)
	GetAllUser(chan *dto.UsersStream)
}
