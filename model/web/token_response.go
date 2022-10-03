package web

type TokenResponse struct {
	Token   string `json:"token,omitempty"`
	Expired int    `json:"expired,omitempty"`
}
