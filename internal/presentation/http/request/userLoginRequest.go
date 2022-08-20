package request

type UserLoginRequest struct {
	Username string `json:"username" param:"username" form:"username"`
	Password string `json:"password" param:"password" form:"password"`
}
