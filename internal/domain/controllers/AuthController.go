package controllers

import (
	"github.com/hamidteimouri/htutils/colog"
	"laramanpurego/internal/domain/repo"
)

type UserController struct {
	repo *repo.UserRepository
}

func NewUserController(repo *repo.UserRepository) *UserController {
	return &UserController{repo: repo}
}

//var tras map[string] locales.Translator = {
//	"fa" :fa.New(),
//}

func (u *UserController) Login(username, password string) error {
	colog.DoBgBlue(username)
	colog.DoBgBlue(password)
	return nil
}

func (u *UserController) Register(name, family, username, password string) error {
	colog.DoBgYellow("name: " + name)
	colog.DoBgYellow(family)
	colog.DoBgYellow(username)
	colog.DoBgYellow(password)
	return nil
}
