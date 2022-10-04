package helper

import (
	"time"

	"github.com/rendyuwu/golang-shortner-link/model/domain"
	"github.com/rendyuwu/golang-shortner-link/model/web"
)

func ToTokenResponse(token domain.Token) web.TokenResponse {
	date := time.Unix(int64(token.Expired), 0)
	loc, err := time.LoadLocation("Asia/Jakarta")

	date = date.In(loc)
	PanicIfError(err)
	return web.TokenResponse{
		Token:   token.Token,
		Expired: date.String(),
	}
}

func ToShortnerResponse(shortner domain.Shortner) web.ShortnerResponse {
	date := time.Unix(int64(shortner.Expired), 0)
	loc, err := time.LoadLocation("Asia/Jakarta")
	PanicIfError(err)

	date = date.In(loc)
	return web.ShortnerResponse{
		Url:        shortner.Url,
		Code:       shortner.Code,
		CustomCode: shortner.CustomCode.String,
		Expired:    date.String(),
	}
}
