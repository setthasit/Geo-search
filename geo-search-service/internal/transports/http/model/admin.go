package model

type SignInRequest struct {
	Username string `json:"username" validate:"required,alphanumunicode"`
	Password string `json:"password" validate:"required"`
}
