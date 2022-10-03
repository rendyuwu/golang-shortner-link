package repository

import (
	"context"
	"database/sql"

	"github.com/rendyuwu/golang-shortner-link/model/domain"
)

type TokenRepository interface {
	Save(ctx context.Context, tx *sql.Tx, token domain.Token) domain.Token
	FindByToken(ctx context.Context, tx *sql.Tx, token string) (domain.Token, error)
}
