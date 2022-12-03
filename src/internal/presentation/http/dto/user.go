package dto

import (
	"goservicetemplate/internal/domain/entity"
	"goservicetemplate/pkg/helpers"
)

type User struct {
	Name      string `json:"name"`
	Family    string `json:"family"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *User) ConvertEntityToDTO(user *entity.User) {
	u.Name = user.Name
	u.Family = user.Family
	u.Email = user.Email
	u.Mobile = user.Mobile
	u.CreatedAt = user.CreatedAt.Format(helpers.TimeFullFormat)
	u.UpdatedAt = user.UpdatedAt.Format(helpers.TimeFullFormat)
}
