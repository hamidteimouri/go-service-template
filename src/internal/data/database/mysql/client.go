package mysql

import (
	"errors"
	"fmt"
	"github.com/hamidteimouri/htutils/htcolog"
	"gorm.io/gorm"
	"laramanpurego/internal/domain/dto"
	"laramanpurego/internal/domain/entity"
	"laramanpurego/pkg/hterror"
	"time"
)

type mysql struct {
	db *gorm.DB
}

func NewMysql(db *gorm.DB) *mysql {
	return &mysql{db: db}
}

func gormErrorToHtError(err error) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return hterror.ErrNotFound
	case gorm.ErrInvalidDB:
		return hterror.ErrorConnection
	default:
		return hterror.ErrorConnection
	}
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
		return nil, gormErrorToHtError(result.Error)
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
	userModel.Password = user.Password

	/* update a column */
	// m.db.Model(&userModel).Update("password", &userModel.Password)

	/* update multiple columns */
	m.db.Model(&userModel).Updates(UserModel{
		Name:     user.Name,
		Family:   user.Family,
		Email:    user.Email,
		Mobile:   user.Mobile,
		Password: user.Password,
	})

	tx := m.db.Save(&userModel)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("skdjfkas jdjskdf %s", tx.Error.Error())
		}
		return nil, tx.Error
	}
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
func (m *mysql) GetAllUser(ch chan *dto.UsersStream) {

	rows, err := m.db.Model(&UserModel{}).Rows()
	defer rows.Close()
	defer close(ch)
	if err != nil {
		ch <- &dto.UsersStream{Error: err}
		return
	}

	for rows.Next() {
		var user UserModel
		// ScanRows scan a row into user
		err := m.db.ScanRows(rows, &user)
		if err != nil {
			ch <- &dto.UsersStream{Error: err}
			return
		}
		us := &entity.User{}
		user.ConvertModelToEntity(us)
		ch <- &dto.UsersStream{User: us}
		time.Sleep(time.Second)
	}
	fmt.Println("done")
}
