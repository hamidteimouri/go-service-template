package data

import (
	"laramanpurego/internal/data/database"
	"laramanpurego/internal/domain/dto"
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
	user, err := u.dbds.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) FindByUsername(username string) (*entity.User, error) {
	return nil, nil
}

func (u *userRepository) FindByEmail(email string) (*entity.User, error) {

	user, err := u.dbds.FindUserByEmail(email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return user, nil
}

func (u *userRepository) Save(user *entity.User) (*entity.User, error) {
	insertUser, err := u.dbds.InsertUser(user)
	if err != nil {
		return nil, err
	}
	return insertUser, err
}

func (u *userRepository) Update(user *entity.User) (*entity.User, error) {
	insertUser, err := u.dbds.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return insertUser, err
}

func (u *userRepository) GetAll(ch chan *dto.UsersStream) {
	u.dbds.GetAllUser(ch)
}
