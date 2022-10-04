package main

import (
	"context"
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
	"github.com/rendyuwu/golang-shortner-link/worker"
)

func main() {
	validate := validator.New()
	env, err := env.NewEnv()
	helper.PanicIfError(err)
	db := app.NewDB(env)
	ctx := context.Background()

	tokenRepository := repository.NewTokenRepository()
	shortnerRepository := repository.NewShortnerRepository()

	tokenService := service.NewTokenService(tokenRepository, db)
	shortnerService := service.NewShortnerService(shortnerRepository, db, validate)

	tokenController := controller.NewTokenController(tokenService)
	shortnerController := controller.NewShortnerController(shortnerService)

	router := router.NewRouter(tokenController, shortnerController)

	server := http.Server{
		Addr:    "0.0.0.0:80",
		Handler: middleware.NewAuthMiddleware(router, tokenService),
	}

	// run cleaner
	go worker.Cleaner(tokenService, shortnerService, ctx)

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
