package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rendyuwu/golang-shortner-link/helper"
	"github.com/rendyuwu/golang-shortner-link/model/web"
	"github.com/rendyuwu/golang-shortner-link/service"
)

type TokenControllerImpl struct {
	TokenService service.TokenService
}

func NewTokenController(tokenService service.TokenService) TokenController {
	return &TokenControllerImpl{
		TokenService: tokenService,
	}
}

func (controller *TokenControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tokenResponse := controller.TokenService.Create(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tokenResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
