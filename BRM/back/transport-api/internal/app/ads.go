package app

import (
	"context"
	"transport-api/internal/model/ads"
)

func (a *appImpl) GetAdById(ctx context.Context, id uint64) (ads.Ad, error) {
	return a.ads.GetAdById(ctx, id)
}

func (a *appImpl) GetAdsList(ctx context.Context, params ads.ListParams) ([]ads.Ad, uint, error) {
	return a.ads.GetAdsList(ctx, params)
}

func (a *appImpl) CreateAd(ctx context.Context, companyId uint64, employeeId uint64, ad ads.Ad) (ads.Ad, error) {
	return a.ads.CreateAd(ctx, companyId, employeeId, ad)
}

func (a *appImpl) UpdateAd(ctx context.Context, companyId uint64, employeeId uint64, adId uint64, upd ads.UpdateAd) (ads.Ad, error) {
	return a.ads.UpdateAd(ctx, companyId, employeeId, adId, upd)
}

func (a *appImpl) DeleteAd(ctx context.Context, companyId uint64, employeeId uint64, adId uint64) error {
	return a.ads.DeleteAd(ctx, companyId, employeeId, adId)
}

func (a *appImpl) CreateResponse(ctx context.Context, companyId uint64, employeeId uint64, adId uint64) (ads.Response, error) {
	return a.ads.CreateResponse(ctx, companyId, employeeId, adId)
}

func (a *appImpl) GetResponses(ctx context.Context, companyId uint64, employeeId uint64, limit uint, offset uint) ([]ads.Response, uint, error) {
	return a.ads.GetResponses(ctx, companyId, employeeId, limit, offset)
}

func (a *appImpl) GetAdsIndustries(ctx context.Context) (map[string]uint64, error) {
	return a.ads.GetIndustries(ctx)
}
