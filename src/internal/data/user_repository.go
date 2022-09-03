package data

import (
	"github.com/hamidteimouri/htutils/colog"
	"laramanpurego/internal/data/database"
	"laramanpurego/internal/domain/entity"
)

type userRepository struct {
	dbds database.DbDatasourceInterface
}

func NewUserRepository(ds database.DbDatasourceInterface) *userRepository {
	return &userRepository{
		dbds: ds,
	}
}

func (u *userRepository) FindById(id string) (*entity.User, error) {
	colog.DoGreen("finding by user id ")
	user, err := u.dbds.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) FindByUsername(username string) (*entity.User, error) {
	colog.DoPurple("find by username is calling")
	return nil, nil
}

func (u *userRepository) FindByEmail(email string) (*entity.User, error) {
	colog.DoPurple("find by email is calling")




	user, err := u.dbds.FindUserByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Save(user *entity.User) (*entity.User, error) {
	return nil, nil
}
