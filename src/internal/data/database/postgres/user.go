package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"goservicetemplate/internal/domain/entity"
	"time"
)

type UserModel struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"size:256"`
	Family    string    `gorm:"size:256"`
	Email     string    `gorm:"size:256"`
	Mobile    string    `gorm:"size:128"`
	Password  string    `gorm:"size:256"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index;default:null"`
}

func (u *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) ConvertEntityToModel(user *entity.User) {
	u.Name = user.Name
	u.Family = user.Family
	u.Email = user.Email
	u.Mobile = user.Mobile
	u.Password = user.Password
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt
}

func (u *UserModel) ConvertModelToEntity(user *entity.User) {
	user.Id = fmt.Sprint(u.Id)
	user.Name = u.Name
	user.Family = u.Family
	user.Email = u.Email
	user.Mobile = u.Mobile
	user.Password = u.Password
	user.CreatedAt = u.CreatedAt
	user.UpdatedAt = u.UpdatedAt
}
