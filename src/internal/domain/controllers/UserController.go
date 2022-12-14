package controllers

import (
	"errors"
	"github.com/sirupsen/logrus"
	"goservicetemplate/internal/domain/dto"
	"goservicetemplate/internal/domain/entity"
	"goservicetemplate/internal/domain/repo"
	"goservicetemplate/pkg/helpers"
)

type UserController struct {
	repo repo.UserRepository
}

func NewUserController(repo repo.UserRepository) *UserController {
	return &UserController{repo: repo}
}

func (u *UserController) Login(username, password string) (token string, err error) {
	user, err := u.repo.FindByEmail(username)
	if err != nil {
		/* there is an error while using database */
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("error while find user by email")
		return "", errors.New("internal error")
	}
	if user == nil {
		/* user not found in database */
		return "", errors.New("incorrect username or password")
	}

	if helpers.HashCheck(password, user.Password) == false {
		/* password is wrong */
		return "", errors.New("incorrect username or password")
	}
	token, err = helpers.JwtGeneration(user.Id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("error while generate jwt token")
		return "", err
	}
	return token, nil
}

func (u *UserController) Register(name, family, username, password string) error {

	user, err := u.repo.FindByEmail(username)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("user exists")
	}

	/* hash password */
	hashed, err := helpers.HashMake(password)
	if err != nil {
		return err
	}
	usr := entity.User{
		Name:     name,
		Family:   family,
		Email:    username,
		Password: hashed,
	}
	_, err = u.repo.Save(&usr)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("error while insert new user")
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

func (u *UserController) ChangePassword(user *entity.User, newPassword string) (ok bool, err error) {
	usr, err := u.repo.FindById(user.Id)
	if err != nil {
		return false, err
	}

	hashed, err := helpers.HashMake(newPassword)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("error while make password hash")
		return false, errors.New("something went wrong")
	}
	usr.Password = hashed
	_, err = u.repo.Update(usr)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("error while update user in database")
		return false, errors.New("something went wrong")
	}

	return true, nil
}

func (u *UserController) Update(user *entity.User) (result *entity.User, ok bool, err error) {

	result, err = u.repo.Update(user)
	if err != nil {
		return nil, false, err
	}

	return result, true, nil
}

func (u *UserController) GetAll() <-chan *dto.UsersStream {
	ch := make(chan *dto.UsersStream)
	go u.repo.GetAll(ch)
	return ch
}
