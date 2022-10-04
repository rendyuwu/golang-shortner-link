package worker

import (
	"context"
	"time"

	"github.com/rendyuwu/golang-shortner-link/service"
)

func Cleaner(tokenService service.TokenService, shortnerService service.ShortnerService, ctx context.Context) {
	for {
		tokenService.Clear(ctx)
		shortnerService.Clear(ctx)

		time.Sleep(1 * time.Second)
	}
}
