package repo

import (
	"laramanpurego/internal/domain/models"
)

type UserRepository interface {
	FindByUsername(username string) (models.User, error)
}
