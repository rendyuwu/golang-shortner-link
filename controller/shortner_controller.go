package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ShortnerController interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindByCode(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Find(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
