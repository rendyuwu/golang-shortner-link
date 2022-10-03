package web

type ShortnerCreateRequest struct {
	Url        string `json:"url,omitempty" validate:"required,max=500,url"`
	CustomCode string `json:"custom_code,omitempty" validate:"omitempty,max=200"`
}
