package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rendyuwu/golang-shortner-link/constanta"
	"github.com/rendyuwu/golang-shortner-link/helper"
	"github.com/rendyuwu/golang-shortner-link/model/domain"
)

type TokenRepositoryImpl struct{}

func NewTokenRepository() TokenRepository {
	return &TokenRepositoryImpl{}
}

func (repository *TokenRepositoryImpl) Save(ctx context.Context, tx *sql.Tx) domain.Token {
	var randomToken string
	for {
		randomToken = helper.RandomToken()
		_, err := repository.FindByToken(ctx, tx, randomToken)
		if err == nil {
			continue
		}
		break
	}

	token := domain.Token{
		Token:   randomToken,
		Expired: int(helper.AddTimestamp(constanta.TOKEN_EXPIRED)),
	}
	SQL := "INSERT INTO token(token, expired) VALUES(?,?)"
	_, err := tx.ExecContext(ctx, SQL, token.Token, token.Expired)
	helper.PanicIfError(err)

	return token
}

func (repository *TokenRepositoryImpl) FindByToken(ctx context.Context, tx *sql.Tx, token string) (domain.Token, error) {
	SQL := "SELECT token, expired FROM token WHERE token = ?"
	rows, err := tx.QueryContext(ctx, SQL, token)
	helper.PanicIfError(err)
	defer rows.Close()

	dataToken := domain.Token{}
	if rows.Next() {
		err := rows.Scan(&dataToken.Token, &dataToken.Expired)
		helper.PanicIfError(err)
		return dataToken, nil
	} else {
		return dataToken, errors.New("token not found")
	}
}
