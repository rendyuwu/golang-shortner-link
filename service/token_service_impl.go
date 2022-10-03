package service

import (
	"context"
	"database/sql"

	"github.com/rendyuwu/golang-shortner-link/constanta"
	"github.com/rendyuwu/golang-shortner-link/exception"
	"github.com/rendyuwu/golang-shortner-link/helper"
	"github.com/rendyuwu/golang-shortner-link/model/domain"
	"github.com/rendyuwu/golang-shortner-link/model/web"
	"github.com/rendyuwu/golang-shortner-link/repository"
)

type TokenServiceImpl struct {
	TokenRepository repository.TokenRepository
	DB              *sql.DB
}

func NewTokenService(tokenRepository repository.TokenRepository, DB *sql.DB) TokenService {
	return &TokenServiceImpl{
		TokenRepository: tokenRepository,
		DB:              DB,
	}
}

func (service *TokenServiceImpl) Create(ctx context.Context) web.TokenResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var randomToken string
	for {
		randomToken = helper.RandomToken(constanta.TOKEN_LENGTH)
		_, err := service.TokenRepository.FindByToken(ctx, tx, randomToken)
		if err == nil {
			continue
		}
		break
	}

	token := domain.Token{
		Token:   randomToken,
		Expired: int(helper.AddTimestamp(constanta.TOKEN_EXPIRED)),
	}

	token = service.TokenRepository.Save(ctx, tx, token)
	return helper.ToTokenResponse(token)
}

func (service *TokenServiceImpl) FindByToken(ctx context.Context, request string) web.TokenResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	token, err := service.TokenRepository.FindByToken(ctx, tx, request)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTokenResponse(token)
}