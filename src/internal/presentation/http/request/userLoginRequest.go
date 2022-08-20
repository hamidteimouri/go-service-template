package request

type UserLoginRequest struct {
	Username string `json:"username" param:"username" form:"username" validate:"required,email"`
	Password string `json:"password" param:"password" form:"password" validate:"required"`
}
