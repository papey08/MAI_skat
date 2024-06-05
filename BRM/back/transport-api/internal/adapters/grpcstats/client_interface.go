package grpcstats

import (
	"context"
	"transport-api/internal/model/stats"
)

//go:generate protoc pb/stats_client.proto --proto_path=pb --go-grpc_out=require_unimplemented_servers=false:. --go_out=.

type StatsClient interface {
	GetCompanyMainPage(ctx context.Context, companyId uint64) (stats.MainPage, error)
}
