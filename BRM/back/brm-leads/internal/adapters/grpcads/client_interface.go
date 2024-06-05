package grpcads

import (
	"brm-leads/internal/model"
	"context"
)

//go:generate protoc pb/ads_client.proto --proto_path=pb --go-grpc_out=require_unimplemented_servers=false:. --go_out=.

type AdsClient interface {
	GetAdData(ctx context.Context, adId uint64) (model.AdData, error)
}
