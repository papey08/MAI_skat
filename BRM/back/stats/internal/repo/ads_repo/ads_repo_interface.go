package ads_repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdsRepo interface {
	GetActiveAdsAmount(ctx context.Context, companyId uint64) (uint, error)
}

func New(pool *pgxpool.Pool) AdsRepo {
	return &adsRepoImpl{
		Pool: pool,
	}
}
