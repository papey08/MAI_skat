package grpcstats

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"transport-api/internal/adapters/grpcstats/pb"
)

type statsClientImpl struct {
	cli pb.StatsServiceClient
}

func NewStatsClient(ctx context.Context, addr string) (StatsClient, error) {
	if conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return &statsClientImpl{}, fmt.Errorf("grpc stats client: %w", err)
	} else {
		return &statsClientImpl{
			cli: pb.NewStatsServiceClient(conn),
		}, nil
	}
}
