package mysql

import (
	"gorm.io/gorm"
	"laramanpurego/internal/domain/entity"
)

type UserModel struct {
	gorm.Model
	Name   string
	Family string
	Email  string
	Mobile string
}

func (u *UserModel) ConvertEntityToModel(user *entity.User) {
	u.ID = user.Id
	u.Name = user.Name
	u.Family = user.Family
	u.Email = user.Email
	u.Mobile = user.Mobile
	u.Mobile = user.Mobile
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt
}
