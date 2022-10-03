package helper

import (
	"time"

	"github.com/rendyuwu/golang-shortner-link/model/domain"
	"github.com/rendyuwu/golang-shortner-link/model/web"
)

func ToTokenResponse(token domain.Token) web.TokenResponse {
	return web.TokenResponse{
		Token:   token.Token,
		Expired: token.Expired,
	}
}

func ToShortnerResponse(shortner domain.Shortner) web.ShortnerResponse {
	date := time.Unix(int64(shortner.Expired), 0)
	loc, err := time.LoadLocation("Asia/Jakarta")
	PanicIfError(err)

	date = date.In(loc)
	return web.ShortnerResponse{
		Url:        shortner.Url,
		CustomCode: shortner.CustomCode.String,
		Expired:    date.String(),
	}
}
