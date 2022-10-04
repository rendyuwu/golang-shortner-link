package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rendyuwu/golang-shortner-link/controller"
	"github.com/rendyuwu/golang-shortner-link/exception"
)

func NewRouter(tokenController controller.TokenController, shortnerController controller.ShortnerController) *httprouter.Router {
	router := httprouter.New()

	// route token
	router.POST("/api/v1/token", tokenController.Create)

	// route shortner
	// router.GET("/:code", shortnerController.Find)

	router.GET("/api/v1/shortner/:code", shortnerController.FindByCode)
	router.POST("/api/v1/shortner", shortnerController.Create)

	router.PanicHandler = exception.ErrorHandler

	return router
}
