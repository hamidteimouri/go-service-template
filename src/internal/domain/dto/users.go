package dto

import "goservicetemplate/internal/domain/entity"

type UsersStream struct {
	Error error
	User  *entity.User
}
