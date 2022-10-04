package repository

import (
	"context"
	"database/sql"

	"github.com/rendyuwu/golang-shortner-link/model/domain"
)

type ShortnerRepository interface {
	Save(ctx context.Context, tx *sql.Tx, shortner domain.Shortner) domain.Shortner
	FindByCode(ctx context.Context, tx *sql.Tx, code string) (domain.Shortner, error)
	FindByCustomCode(ctx context.Context, tx *sql.Tx, code string) (domain.Shortner, error)
	Clear(ctx context.Context, tx *sql.Tx, timestamp int)
}
