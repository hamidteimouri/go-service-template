package request

type UserUpdateRequest struct {
	Name   string `json:"name" form:"name" param:"name" validate:"required"`
	Family string `json:"family" form:"family" param:"family" validate:"required"`
	Mobile string `json:"mobile" form:"mobile" param:"mobile"`
}
