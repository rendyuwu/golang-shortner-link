package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rendyuwu/golang-shortner-link/helper"
	"github.com/rendyuwu/golang-shortner-link/model/web"
	"github.com/rendyuwu/golang-shortner-link/service"
)

type ShortnerControllerImpl struct {
	ShortnerService service.ShortnerService
}

func NewShortnerController(shortnerService service.ShortnerService) ShortnerController {
	return &ShortnerControllerImpl{
		ShortnerService: shortnerService,
	}
}

func (controller *ShortnerControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	shortnerCreateRequest := web.ShortnerCreateRequest{}
	helper.ReadFromRequestBody(r, &shortnerCreateRequest)

	shortnerResponse := controller.ShortnerService.Create(r.Context(), shortnerCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   shortnerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *ShortnerControllerImpl) FindByCode(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	code := p.ByName("code")

	shortnerResponse := controller.ShortnerService.FindByCode(r.Context(), code)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   shortnerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *ShortnerControllerImpl) Find(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	code := p.ByName("code")

	shortnerResponse := controller.ShortnerService.FindByCode(r.Context(), code)
	http.Redirect(w, r, shortnerResponse.Url, http.StatusPermanentRedirect)
}
