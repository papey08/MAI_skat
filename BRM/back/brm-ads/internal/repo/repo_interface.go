package repo

import (
	"brm-ads/internal/model"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdRepo interface {
	GetAdById(ctx context.Context, id uint64) (model.Ad, error)
	GetAdsList(ctx context.Context, params model.AdsListParams) ([]model.Ad, uint, error)
	CreateAd(ctx context.Context, ad model.Ad) (model.Ad, error)
	UpdateAd(ctx context.Context, adId uint64, upd model.UpdateAd) (model.Ad, error)
	DeleteAd(ctx context.Context, adId uint64) error

	CreateResponse(ctx context.Context, resp model.Response) (model.Response, error)
	GetResponses(ctx context.Context, companyId uint64, limit uint, offset uint) ([]model.Response, uint, error)

	GetIndustries(ctx context.Context) (map[string]uint64, error)
}

func New(pool *pgxpool.Pool) AdRepo {
	return &adRepoImpl{
		Pool: pool,
	}
}
