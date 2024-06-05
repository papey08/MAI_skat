package grpcads

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"transport-api/internal/adapters/grpcads/pb"
)

type adsClientImpl struct {
	cli pb.AdsServiceClient
}

func NewAdsClient(ctx context.Context, addr string) (AdsClient, error) {
	if conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return &adsClientImpl{}, fmt.Errorf("grpc ads client: %w", err)
	} else {
		return &adsClientImpl{
			cli: pb.NewAdsServiceClient(conn),
		}, nil
	}
}
