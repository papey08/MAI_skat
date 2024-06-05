package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"images/internal/model"
)

type ImageRepo interface {
	AddImage(ctx context.Context, img model.Image) (uint64, error)
	GetImage(ctx context.Context, id uint64) (model.Image, error)
}

func New(pool *pgxpool.Pool) ImageRepo {
	return &repoImpl{
		Pool: pool,
	}
}
