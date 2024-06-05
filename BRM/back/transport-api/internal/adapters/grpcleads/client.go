package grpcleads

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"transport-api/internal/adapters/grpcleads/pb"
)

type leadsClientImpl struct {
	cli pb.LeadsServiceClient
}

func NewLeadsClient(ctx context.Context, addr string) (LeadsClient, error) {
	if conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return &leadsClientImpl{}, fmt.Errorf("grpc leads client: %w", err)
	} else {
		return &leadsClientImpl{
			cli: pb.NewLeadsServiceClient(conn),
		}, nil
	}
}
