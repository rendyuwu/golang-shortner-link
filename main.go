package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rendyuwu/golang-shortner-link/app"
	"github.com/rendyuwu/golang-shortner-link/controller"
	"github.com/rendyuwu/golang-shortner-link/env"
	"github.com/rendyuwu/golang-shortner-link/helper"
	"github.com/rendyuwu/golang-shortner-link/middleware"
	"github.com/rendyuwu/golang-shortner-link/repository"
	"github.com/rendyuwu/golang-shortner-link/router"
	"github.com/rendyuwu/golang-shortner-link/service"
)

func main() {
	validate := validator.New()
	env, err := env.NewEnv()
	helper.PanicIfError(err)
	db := app.NewDB(env)

	tokenRepository := repository.NewTokenRepository()
	shortnerRepository := repository.NewShortnerRepository()

	tokenService := service.NewTokenService(tokenRepository, db)
	shortnerService := service.NewShortnerService(shortnerRepository, db, validate)

	tokenController := controller.NewTokenController(tokenService)
	shortnerController := controller.NewShortnerController(shortnerService)

	router := router.NewRouter(tokenController, shortnerController)

	server := http.Server{
		Addr:    "0.0.0.0:8888",
		Handler: middleware.NewAuthMiddleware(router, tokenService),
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
