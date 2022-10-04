package service

import (
	"context"

	"github.com/rendyuwu/golang-shortner-link/model/web"
)

type ShortnerService interface {
	Create(ctx context.Context, request web.ShortnerCreateRequest) web.ShortnerResponse
	FindByCode(ctx context.Context, request string) web.ShortnerResponse
	Clear(ctx context.Context)
}
