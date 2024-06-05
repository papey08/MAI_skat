package grpcnotifications

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"transport-api/internal/adapters/grpcnotifications/pb"
)

type notificationsClientImpl struct {
	cli pb.NotificationsServiceClient
}

func NewNotificationsClient(ctx context.Context, addr string) (NotificationsClient, error) {
	if conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return &notificationsClientImpl{}, fmt.Errorf("grpc notifications client: %w", err)
	} else {
		return &notificationsClientImpl{
			cli: pb.NewNotificationsServiceClient(conn),
		}, nil
	}
}
