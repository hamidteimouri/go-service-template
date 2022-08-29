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

func (u *userRepository) FindByUsername(username string) (*entity.User, error) {
	colog.DoPurple("I am here")
	return nil, nil
}

func (u *userRepository) Save(user *entity.User) (*entity.User, error) {
	panic("implement me")
}
