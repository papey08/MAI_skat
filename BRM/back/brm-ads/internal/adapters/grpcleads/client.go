package grpcleads

import (
	"brm-ads/internal/adapters/grpcleads/pb"
	"brm-ads/internal/model"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type leadsClientImpl struct {
	cli pb.LeadsServiceClient
}

func NewLeadsClient(ctx context.Context, addr string) (LeadsClient, error) {
	if conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return &leadsClientImpl{}, fmt.Errorf("grpc leads client: %w", err)
	} else {
		return &leadsClientImpl{cli: pb.NewLeadsServiceClient(conn)}, nil
	}
}

func (l *leadsClientImpl) CreateLead(ctx context.Context, adId uint64, clientCompany uint64, clientEmployee uint64) error {
	_, err := l.cli.CreateLead(ctx, &pb.CreateLeadRequest{
		AdId:           adId,
		ClientCompany:  clientCompany,
		ClientEmployee: clientEmployee,
	})
	if err != nil {
		return model.ErrLeadsServiceError
	}
	return nil
}
