package grpcauth

import (
	"brm-core/internal/model"
	"context"
)

//go:generate protoc pb/auth_client.proto --proto_path=pb --go-grpc_out=require_unimplemented_servers=false:. --go_out=.

type AuthClient interface {
	RegisterEmployee(ctx context.Context, creds model.EmployeeCredentials) error
	DeleteEmployee(ctx context.Context, email string) error
}
