package dto

import "laramanpurego/internal/domain/entity"

type UsersStream struct {
	Error error
	User  *entity.User
}
