package grpccore

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"transport-api/internal/adapters/grpccore/pb"
)

type coreClientImpl struct {
	cli pb.CoreServiceClient
}

func NewCoreClient(ctx context.Context, addr string) (CoreClient, error) {
	if conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return &coreClientImpl{}, fmt.Errorf("grpc core client: %w", err)
	} else {
		return &coreClientImpl{
			cli: pb.NewCoreServiceClient(conn),
		}, nil
	}
}
