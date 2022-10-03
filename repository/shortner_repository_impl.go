package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rendyuwu/golang-shortner-link/constanta"
	"github.com/rendyuwu/golang-shortner-link/helper"
	"github.com/rendyuwu/golang-shortner-link/model/domain"
)

type ShortnerRepositoryImpl struct{}

func NewSHortnerRepository() ShortnerRepository {
	return &ShortnerRepositoryImpl{}
}

func (repository *ShortnerRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, shortner domain.Shortner) domain.Shortner {
	var randomCode string
	for {
		randomCode = helper.RandomToken(constanta.SHORTNER_LENGTH)
		_, err := repository.FindByCode(ctx, tx, randomCode)
		if err == nil {
			continue
		}
		break
	}

	SQL := "INSERT INTO shortner(code, custom_code, url, expired) VALUES(?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, shortner.Code, shortner.CustomCode, shortner.Url, shortner.Expired)
	helper.PanicIfError(err)

	return shortner
}

func (repository *ShortnerRepositoryImpl) FindByCode(ctx context.Context, tx *sql.Tx, code string) (domain.Shortner, error) {
	SQL := "SELECT code, custom_code, url, expired FROM shortner WHERE code = ? OR custom_code = ?"
	rows, err := tx.QueryContext(ctx, SQL, code, code)
	helper.PanicIfError(err)
	defer rows.Close()

	shortner := domain.Shortner{}
	if rows.Next() {
		err := rows.Scan(&shortner.Code, &shortner.CustomCode, &shortner.Url, &shortner.Expired)
		helper.PanicIfError(err)
		return shortner, nil
	} else {
		return shortner, errors.New("shortner link not found")
	}
}
