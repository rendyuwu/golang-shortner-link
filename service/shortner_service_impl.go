package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rendyuwu/golang-shortner-link/constanta"
	"github.com/rendyuwu/golang-shortner-link/exception"
	"github.com/rendyuwu/golang-shortner-link/helper"
	"github.com/rendyuwu/golang-shortner-link/model/domain"
	"github.com/rendyuwu/golang-shortner-link/model/web"
	"github.com/rendyuwu/golang-shortner-link/repository"
)

type ShortnerServiceImpl struct {
	ShortnerRepository repository.ShortnerRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewShortnerService(shortnerRepository repository.ShortnerRepository, DB *sql.DB, validate *validator.Validate) ShortnerService {
	return &ShortnerServiceImpl{
		ShortnerRepository: shortnerRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *ShortnerServiceImpl) Create(ctx context.Context, request web.ShortnerCreateRequest) web.ShortnerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// generate random code
	var randomCode string
	for {
		randomCode = helper.RandomToken(constanta.SHORTNER_LENGTH)
		_, err := service.ShortnerRepository.FindByCode(ctx, tx, randomCode)
		if err == nil {
			continue
		}
		break
	}

	// check if custom code exist
	shortner, err := service.ShortnerRepository.FindByCustomCode(ctx, tx, request.CustomCode)
	if err == nil {
		panic(exception.NewAlreadyExistError("custom code already exist"))
	}

	shortner = domain.Shortner{
		Code:       randomCode,
		CustomCode: sql.NullString{String: request.CustomCode, Valid: true},
		Url:        request.Url,
		Expired:    int(helper.AddTimestamp(constanta.SHORTNER_EXPIRED)),
	}

	shortner = service.ShortnerRepository.Save(ctx, tx, shortner)

	return helper.ToShortnerResponse(shortner)
}

func (service *ShortnerServiceImpl) FindByCode(ctx context.Context, request string) web.ShortnerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	found := false
	shortner, err := service.ShortnerRepository.FindByCustomCode(ctx, tx, request)
	if err == nil {
		found = true
	}

	if found {
		return helper.ToShortnerResponse(shortner)
	} else {
		shortner, err := service.ShortnerRepository.FindByCode(ctx, tx, request)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}
		return helper.ToShortnerResponse(shortner)
	}
}

func (service *ShortnerServiceImpl) Clear(ctx context.Context) {
	timestamp := time.Now().Unix()

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.ShortnerRepository.Clear(ctx, tx, int(timestamp))
}
