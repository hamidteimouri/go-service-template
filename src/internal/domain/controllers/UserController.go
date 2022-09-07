package controllers

import (
	"github.com/hamidteimouri/htutils/colog"
	"laramanpurego/internal/domain/entity"
	"laramanpurego/internal/domain/repo"
	"laramanpurego/pkg/helpers"
)

type UserController struct {
	repo repo.UserRepository
}

func NewUserController(repo repo.UserRepository) *UserController {
	return &UserController{repo: repo}
}

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
		colog.DoRed(err.Error())
		return "", err
	}
	colog.DoGreen(token)
	return token, nil
}

func (u *UserController) Register(name, family, username, password string) error {
	colog.DoBgGreen("register method called in controller")
	hashed, err := helpers.HashMake(password)
	if err != nil {
		return err
	}
	user := entity.User{
		Name:     name,
		Family:   family,
		Email:    username,
		Password: hashed,
	}
	_, err = u.repo.Save(&user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserController) GetUserByEmail(email string) (user *entity.User, err error) {
	user, err = u.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserController) GetUserByID(id string) (user *entity.User, err error) {
	user, err = u.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
