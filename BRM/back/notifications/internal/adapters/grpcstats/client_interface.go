package grpcstats

import "context"

//go:generate protoc pb/stats_client.proto --proto_path=pb --go-grpc_out=require_unimplemented_servers=false:. --go_out=.

type StatsClient interface {
	SubmitClosedLead(ctx context.Context, producerCompanyId uint64, submit bool) error
}
