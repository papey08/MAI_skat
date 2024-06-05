package grpccore

import "context"

//go:generate protoc pb/core_client.proto --proto_path=pb --go-grpc_out=require_unimplemented_servers=false:. --go_out=.

type CoreClient interface {
	GetCompanyName(ctx context.Context, id uint64) (string, error)
	GetEmployeeById(ctx context.Context, companyId uint64, employeeId uint64, employeeIdToFind uint64) (uint64, uint64, error)
}
