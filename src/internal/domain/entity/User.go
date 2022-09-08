package entity

import (
	"strconv"
	"strings"
	"time"
)

type User struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Family    string    `json:"family"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(id uint, name string, family string, email string, mobile string, password string, createdAt time.Time, updatedAt time.Time) *User {
	return &User{Id: id, Name: name, Family: family, Email: email, Mobile: mobile, Password: password, CreatedAt: createdAt, UpdatedAt: updatedAt}
}

func (u *User) SetEmail(email string) {
	u.Email = strings.ToLower(email)
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetFamily(family string) {
	u.Family = family
}

func (u *User) GetFullName() string {
	return u.Name + " " + u.Family
}
func (u *User) GetIdString() string {
	return strconv.FormatUint(uint64(u.Id), 10)
}
