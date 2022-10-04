package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rendyuwu/golang-shortner-link/helper"
	"github.com/rendyuwu/golang-shortner-link/model/domain"
)

type ShortnerRepositoryImpl struct{}

func NewShortnerRepository() ShortnerRepository {
	return &ShortnerRepositoryImpl{}
}

func (repository *ShortnerRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, shortner domain.Shortner) domain.Shortner {
	SQL := "INSERT INTO shortner(code, custom_code, url, expired) VALUES(?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, shortner.Code, shortner.CustomCode, shortner.Url, shortner.Expired)
	helper.PanicIfError(err)

	return shortner
}

func (repository *ShortnerRepositoryImpl) FindByCode(ctx context.Context, tx *sql.Tx, code string) (domain.Shortner, error) {
	SQL := "SELECT code, custom_code, url, expired FROM shortner WHERE code = ?"
	rows, err := tx.QueryContext(ctx, SQL, code)
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

func (repository *ShortnerRepositoryImpl) FindByCustomCode(ctx context.Context, tx *sql.Tx, code string) (domain.Shortner, error) {
	SQL := "SELECT code, custom_code, url, expired FROM shortner WHERE custom_code = ?"
	rows, err := tx.QueryContext(ctx, SQL, code)
	helper.PanicIfError(err)
	defer rows.Close()

	shortner := domain.Shortner{}
	if rows.Next() {
		err := rows.Scan(&shortner.Code, &shortner.CustomCode, &shortner.Url, &shortner.Expired)
		helper.PanicIfError(err)
		return shortner, nil
	} else {
		return shortner, errors.New("custom code already exist")
	}
}
