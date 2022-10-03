package web

type ShortnerResponse struct {
	Url        string `json:"url,omitempty"`
	CustomCode string `json:"custom_code,omitempty"`
	Expired    string `json:"expired,omitempty"`
}
