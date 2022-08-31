package repo

import (
	"laramanpurego/internal/domain/entity"
)

type UserRepository interface {
	FindByUsername(username string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindById(id string) (*entity.User, error)
	Save(user *entity.User) (*entity.User, error)
}
