package grpcleads

import "context"

//go:generate protoc pb/leads_client.proto --proto_path=pb --go-grpc_out=require_unimplemented_servers=false:. --go_out=.

type LeadsClient interface {
	CreateLead(ctx context.Context, adId uint64, clientCompany uint64, clientEmployee uint64) error
}
