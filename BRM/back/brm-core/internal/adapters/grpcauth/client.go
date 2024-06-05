package grpcauth

import (
	"brm-core/internal/adapters/grpcauth/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authClientImpl struct {
	cli pb.AuthServiceClient
}

func NewAuthClient(ctx context.Context, addr string) (AuthClient, error) {
	if conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return &authClientImpl{}, fmt.Errorf("grpc core client: %w", err)
	} else {
		return &authClientImpl{
			cli: pb.NewAuthServiceClient(conn),
		}, nil
	}
}
