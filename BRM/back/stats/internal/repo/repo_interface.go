package repo

import (
	"context"
	"stats/internal/model"
	adsRepo "stats/internal/repo/ads_repo"
	"stats/internal/repo/cache"
	coreRepo "stats/internal/repo/core_repo"
	leadsRepo "stats/internal/repo/leads_repo"
)

type Repo interface {
	GetCompanyMainPageStats(ctx context.Context, id uint64) (model.MainPageStats, error)

	GetCompanyRating(ctx context.Context, id uint64) (float64, error)
	SetCompanyRating(ctx context.Context, id uint64, rating float64) error
}

func New(ads adsRepo.AdsRepo, core coreRepo.CoreRepo, leads leadsRepo.LeadsRepo, cache cache.Cache) Repo {
	return &repoImpl{
		ads:   ads,
		leads: leads,
		core:  core,
		cache: cache,
	}
}
