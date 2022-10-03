package web

type TokenCreateRequest struct {
	Token  string `json:"token,omitempty" validate:"required,alphanum,max=60"`
	Expire int    `json:"expire,omitempty" validate:"required,number"`
}
