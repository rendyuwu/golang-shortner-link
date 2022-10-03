package helper

import (
	"github.com/rendyuwu/golang-shortner-link/model/domain"
	"github.com/rendyuwu/golang-shortner-link/model/web"
)

func ToTokenResponse(token domain.Token) web.TokenResponse {
	return web.TokenResponse{
		Token:   token.Token,
		Expired: token.Expired,
	}
}
