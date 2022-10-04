package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TokenController interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
