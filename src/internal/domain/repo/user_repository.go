package repo

import (
	"laramanpurego/src/internal/domain/models"
)

type UserRepository interface {
	FindByUsername(username string) (models.User, error)
}


