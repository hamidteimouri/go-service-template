package request

type UserChangePasswordRequest struct {
	CurrentPassword      string `json:"current_password" param:"current_password" form:"current_password" validate:"required"`
	Password             string `json:"password" param:"password" form:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" param:"password_confirmation" form:"password_confirmation" validate:"required"`
}
