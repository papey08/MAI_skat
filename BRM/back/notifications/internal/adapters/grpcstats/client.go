package grpcstats

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"notifications/internal/adapters/grpcstats/pb"
	"notifications/internal/model"
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

func (s *statsClientImpl) SubmitClosedLead(ctx context.Context, producerCompanyId uint64, submit bool) error {
	_, err := s.cli.SubmitClosedLead(ctx, &pb.SubmitClosedLeadRequest{
		CompanyId: producerCompanyId,
		Submit:    submit,
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return model.ErrCompanyNotFound
		default:
			return model.ErrStatsError
		}
	}
	return nil
}
