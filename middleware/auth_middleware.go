package middleware

import (
	"net/http"

	"github.com/rendyuwu/golang-shortner-link/helper"
	"github.com/rendyuwu/golang-shortner-link/model/web"
	"github.com/rendyuwu/golang-shortner-link/service"
)

type AuthMiddleware struct {
	Handler      http.Handler
	TokenService service.TokenService
}

func NewAuthMiddleware(handler http.Handler, tokenService service.TokenService) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler, TokenService: tokenService}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := middleware.TokenService.FindByToken(r.Context(), r.Header.Get("X-API-KEY"))

	if err == nil {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
