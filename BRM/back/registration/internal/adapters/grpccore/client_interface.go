package grpccore

import (
	"context"
	"registration/internal/model"
)

//go:generate protoc pb/core_client.proto --proto_path=pb --go-grpc_out=require_unimplemented_servers=false:. --go_out=.

type CoreClient interface {
	CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error)
	GetIndustriesList(ctx context.Context) (map[string]uint64, error)
}
