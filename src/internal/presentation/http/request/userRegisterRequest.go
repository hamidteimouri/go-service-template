package request

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Family   string `json:"family"`
	Username string `json:"username"`
	Password string `json:"password"`
}
