package mysql

import (
	"goservicetemplate/internal/domain/entity"
	"time"
)

type UserModel struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:256"`
	Family    string `gorm:"size:256"`
	Email     string `gorm:"size:256"`
	Mobile    string `gorm:"size:128"`
	Password  string `gorm:"size:256"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index;default:null"`
}

func (u *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) ConvertEntityToModel(user *entity.User) {
	u.ID = user.Id
	u.Name = user.Name
	u.Family = user.Family
	u.Email = user.Email
	u.Mobile = user.Mobile
	u.Password = user.Password
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt
}

func (u *UserModel) ConvertModelToEntity(user *entity.User) {
	user.Id = u.ID
	user.Name = u.Name
	user.Family = u.Family
	user.Email = u.Email
	user.Mobile = u.Mobile
	user.Password = u.Password
	user.CreatedAt = u.CreatedAt
	user.UpdatedAt = u.UpdatedAt
}
