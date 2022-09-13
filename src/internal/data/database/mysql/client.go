package mysql

import (
	"errors"
	"github.com/hamidteimouri/htutils/htcolog"
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

	um := UserModel{}
	user := &entity.User{}

	result := m.db.Table("users").Where("id = ?", id).First(&um)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			htcolog.DoBgBrightRed(id + " Not found")
			return nil, nil
		}
		return nil, result.Error
	}
	um.ConvertModelToEntity(user)
	return user, nil
}

func (m *mysql) FindUserByEmail(email string) (*entity.User, error) {
	um := UserModel{}
	user := &entity.User{}
	result := m.db.Table("users").Where("email = ?", email).First(&um)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	um.ConvertModelToEntity(user)
	return user, nil
}

func (m *mysql) FindUserByMobile(mobile string) (user *entity.User, err error) {
	um := UserModel{}
	user = &entity.User{}

	result := m.db.Model(&UserModel{}).Where("mobile = ?", mobile).Find(&um)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			htcolog.DoBgBrightRed(mobile + " Not found")
			return nil, nil
		}
		return nil, result.Error
	}
	um.ConvertModelToEntity(user)
	return
}

func (m *mysql) UpdateUser(user *entity.User) (*entity.User, error) {
	userModel := UserModel{}
	userModel.ConvertEntityToModel(user)
	//userModel.Password = user.Password

	m.db.Model(&userModel).Update("password", &userModel.Password)

	m.db.Save(&userModel)
	userModel.ConvertModelToEntity(user)
	return user, nil
}

func (m *mysql) InsertUser(user *entity.User) (*entity.User, error) {
	userModel := UserModel{}
	userModel.ConvertEntityToModel(user)

	result := m.db.Create(&userModel)
	if result.Error != nil {
		htcolog.DoRed("err while insert new user : " + result.Error.Error())
		return nil, result.Error
	}
	userModel.ConvertModelToEntity(user)
	return user, nil
}
