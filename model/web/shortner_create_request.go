package web

import "database/sql"

type ShortnerCreateRequest struct {
	Url        string         `json:"url,omitempty" validate:"required,max=500,url"`
	CustomCode sql.NullString `json:"custom_code,omitempty" validate:"omitempty,max=200"`
}
