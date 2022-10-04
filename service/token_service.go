package service

import (
	"context"

	"github.com/rendyuwu/golang-shortner-link/model/web"
)

type TokenService interface {
	Create(ctx context.Context) web.TokenResponse
	FindByToken(ctx context.Context, request string) (web.TokenResponse, error)
	Clear(ctx context.Context)
}
