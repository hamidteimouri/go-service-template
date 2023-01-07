package data

import (
	"goservicetemplate/internal/data/database"
	"goservicetemplate/internal/domain/dto"
	"goservicetemplate/internal/domain/entity"
)

type UserRepository struct {
	dbds database.DbDatasourceInterface
}

func NewUserRepository(ds database.DbDatasourceInterface) *UserRepository {
	return &UserRepository{
		dbds: ds,
	}
}

func (u *UserRepository) FindById(id string) (*entity.User, error) {
	user, err := u.dbds.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) FindByUsername(username string) (*entity.User, error) {
	return nil, nil
}

func (u *UserRepository) FindByEmail(email string) (*entity.User, error) {

	user, err := u.dbds.FindUserByEmail(email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return user, nil
}

func (u *UserRepository) Save(user *entity.User) (*entity.User, error) {
	insertUser, err := u.dbds.InsertUser(user)
	if err != nil {
		return nil, err
	}
	return insertUser, err
}

func (u *UserRepository) Update(user *entity.User) (*entity.User, error) {
	insertUser, err := u.dbds.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return insertUser, err
}

func (u *UserRepository) GetAll(ch chan *dto.UsersStream) {
	u.dbds.GetAllUser(ch)
}
