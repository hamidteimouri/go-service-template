package data

import "laramanpurego/internal/domain/entity"

type UserInterface interface {
	Save(user *entity.User) *entity.User
	FindByEmail(email string) *entity.User
}

type userRepository struct {
}

func newUserRepository() *userRepository {
	return &userRepository{}
}

func (u userRepository) Save(user *entity.User) *entity.User {
	panic("implement me")
}

func (u userRepository) FindByEmail(email string) *entity.User {
	panic("implement me")
}
