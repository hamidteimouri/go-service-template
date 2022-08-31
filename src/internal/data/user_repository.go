package data

import (
	"github.com/hamidteimouri/htutils/colog"
	"laramanpurego/internal/domain/entity"
)

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (u *userRepository) FindById(id string) (*entity.User, error) {
	colog.DoGreen("finding by user id ")
	return nil, nil
}

func (u *userRepository) FindByUsername(username string) (*entity.User, error) {
	colog.DoPurple("find by username is calling")
	return nil, nil
}

func (u *userRepository) FindByEmail(email string) (*entity.User, error) {
	colog.DoPurple("find by email is calling")

	user := entity.User{
		Id:    1,
		Email: "hamid@test.com",
	}
	return &user, nil
}

func (u *userRepository) Save(user *entity.User) (*entity.User, error) {
	return nil, nil
}
