package models

type User struct {
	Id       string
	Name     string
	Family   string
	Email    string
	Password string `json:"-"`
}
