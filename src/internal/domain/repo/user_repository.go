package repo

import (
	"laramanpurego/internal/domain/entity"
)

type UserRepository interface {
	FindByUsername(username string) (entity.User, error)
}
