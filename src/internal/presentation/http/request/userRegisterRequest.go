package request

type UserRegisterRequest struct {
	Name     string `json:"name" param:"name" form:"name" validate:"required"`
	Family   string `json:"family" param:"family" form:"family" validate:"required"`
	Username string `json:"username" param:"username" form:"username" validate:"required,email"`
	Password string `json:"password" param:"password" form:"password" validate:"required"`
}
