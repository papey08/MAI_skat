package repo

import (
	"context"
	"stats/internal/model"
	adsRepo "stats/internal/repo/ads_repo"
	"stats/internal/repo/cache"
	coreRepo "stats/internal/repo/core_repo"
	leadsRepo "stats/internal/repo/leads_repo"
)

type repoImpl struct {
	ads   adsRepo.AdsRepo
	leads leadsRepo.LeadsRepo
	core  coreRepo.CoreRepo

	cache cache.Cache
}

func (r *repoImpl) GetCompanyMainPageStats(ctx context.Context, id uint64) (model.MainPageStats, error) {
	if data, ok := r.cache.GetCompanyData(ctx, id); ok {
		return data, nil
	}

	data, err := r.getStatsFromPermanent(ctx, id)
	if err != nil {
		return model.MainPageStats{}, err
	}

	r.cache.AddCompanyData(ctx, id, data)
	return data, nil
}

func (r *repoImpl) GetCompanyRating(ctx context.Context, id uint64) (float64, error) {
	return r.core.GetCompanyAbsoluteRating(ctx, id)
}

func (r *repoImpl) SetCompanyRating(ctx context.Context, id uint64, rating float64) error {
	return r.core.SetCompanyRating(ctx, id, rating)
}

func (r *repoImpl) getStatsFromPermanent(ctx context.Context, id uint64) (model.MainPageStats, error) {
	data, err := r.leads.GetMainPageLeadsStats(ctx, id)
	if err != nil {
		return model.MainPageStats{}, err
	}

	data.ActiveAdsAmount, err = r.ads.GetActiveAdsAmount(ctx, id)
	if err != nil {
		return model.MainPageStats{}, err
	}

	data.CompanyAbsoluteRating, err = r.core.GetCompanyAbsoluteRating(ctx, id)
	if err != nil {
		return model.MainPageStats{}, err
	}

	data.CompanyRelativeRating, err = r.core.GetCompanyRelativeRating(ctx, id)
	if err != nil {
		return model.MainPageStats{}, err
	}

	return data, nil
}
