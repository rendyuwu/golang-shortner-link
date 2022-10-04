package web

type TokenResponse struct {
	Token   string `json:"token,omitempty"`
	Expired string `json:"expired,omitempty"`
}
