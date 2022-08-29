package controllers

import (
	"github.com/hamidteimouri/htutils/colog"
	"laramanpurego/internal/domain/repo"
	"laramanpurego/pkg/helpers"
)

type UserController struct {
	repo repo.UserRepository
}

func NewUserController(repo repo.UserRepository) *UserController {
	return &UserController{repo: repo}
}

//var tras map[string] locales.Translator = {
//	"fa" :fa.New(),
//}

func (u *UserController) Login(username, password string) (token string, err error) {
	colog.DoBgBlue("login method called")
	_, err = u.repo.FindByUsername(username)
	if err != nil {
		colog.DoRed(err.Error())
		return "", err
	}

	colog.DoBgBlue(username)
	colog.DoBgBlue(password)
	token, err = helpers.JwtGeneration("حمید", "تیموری", username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *UserController) Register(name, family, username, password string) error {
	colog.DoBgYellow("name: " + name)
	colog.DoBgYellow(family)
	colog.DoBgYellow(username)
	colog.DoBgYellow(password)
	return nil
}
