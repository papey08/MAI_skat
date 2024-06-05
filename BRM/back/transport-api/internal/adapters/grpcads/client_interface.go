package grpcads

import (
	"context"
	"transport-api/internal/model/ads"
)

//go:generate protoc pb/ads_client.proto --proto_path=pb --go-grpc_out=require_unimplemented_servers=false:. --go_out=.

type AdsClient interface {
	GetAdById(ctx context.Context, id uint64) (ads.Ad, error)
	GetAdsList(ctx context.Context, params ads.ListParams) ([]ads.Ad, uint, error)
	CreateAd(ctx context.Context, companyId uint64, employeeId uint64, ad ads.Ad) (ads.Ad, error)
	UpdateAd(ctx context.Context, companyId uint64, employeeId uint64, adId uint64, upd ads.UpdateAd) (ads.Ad, error)
	DeleteAd(ctx context.Context, companyId uint64, employeeId uint64, adId uint64) error

	CreateResponse(ctx context.Context, companyId uint64, employeeId uint64, adId uint64) (ads.Response, error)
	GetResponses(ctx context.Context, companyId uint64, employeeId uint64, limit uint, offset uint) ([]ads.Response, uint, error)

	GetIndustries(ctx context.Context) (map[string]uint64, error)
}
