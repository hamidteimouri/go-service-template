package controllers

import (
	"errors"
	"fmt"
	"github.com/hamidteimouri/htutils/colog"
	"laramanpurego/internal/domain/entity"
	"laramanpurego/internal/domain/repo"
	"laramanpurego/pkg/helpers"
	"strconv"
)

type UserController struct {
	repo repo.UserRepository
}

func NewUserController(repo repo.UserRepository) *UserController {
	return &UserController{repo: repo}
}

func (u *UserController) Login(username, password string) (token string, err error) {
	colog.DoBgBlue("login method called")
	user, err := u.repo.FindByEmail(username)
	if err != nil {
		colog.DoRed(err.Error())
		return "", errors.New("incorrect username or password")
	}
	fmt.Println(user)

	if helpers.HashCheck(password, user.Password) == false {
		return "", errors.New("incorrect username or password")
	}
	var userId string = strconv.FormatUint(uint64(user.Id), 10)
	token, err = helpers.JwtGeneration(userId)
	if err != nil {
		colog.DoRed(err.Error())
		return "", err
	}
	colog.DoGreen(token)
	return token, nil
}

func (u *UserController) Register(name, family, username, password string) error {
	colog.DoBgGreen("register method called in controller")
	/* hash password */
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
