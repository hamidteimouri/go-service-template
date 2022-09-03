package mysql

import (
	"errors"
	"github.com/hamidteimouri/htutils/colog"
	"gorm.io/gorm"
	"laramanpurego/internal/domain/entity"
)

type mysql struct {
	db *gorm.DB
}

func NewMysql(db *gorm.DB) *mysql {
	return &mysql{db: db}
}

func (m *mysql) FindUserById(id string) (*entity.User, error) {
	colog.DoGreen("mysql find user by id")
	panic("implement me")
}

func (m *mysql) FindUserByEmail(email string) (*entity.User, error) {
	colog.DoGreen("mysql find user by email")

	um := UserModel{}
	um.Email = email
	user := &entity.User{}

	result := m.db.Model(&UserModel{}).Where("email = ?", email).Find(&um)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	um.ConvertEntityToModel(user)
	return user, nil

}

func (m *mysql) UpdateUser(user *entity.User) (*entity.User, error) {
	userModel := UserModel{}
	userModel.ConvertEntityToModel(user)
	panic("implement me")
}

func (m *mysql) InsertUser(user *entity.User) (*entity.User, error) {
	userModel := UserModel{}
	userModel.ConvertEntityToModel(user)
	panic("implement me")
}
