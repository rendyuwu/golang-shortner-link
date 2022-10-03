package domain

import "database/sql"

type Shortner struct {
	Id         int
	Code       string
	CustomCode sql.NullString
	Url        string
}
