package grpcstats

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"transport-api/internal/adapters/grpcstats/pb"
	"transport-api/internal/model"
	"transport-api/internal/model/stats"
)

func (s *statsClientImpl) GetCompanyMainPage(ctx context.Context, companyId uint64) (stats.MainPage, error) {
	resp, err := s.cli.GetCompanyMainPage(ctx, &pb.GetCompanyMainPageRequest{
		CompanyId: companyId,
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return stats.MainPage{}, model.ErrCompanyNotExists
		case codes.ResourceExhausted:
			return stats.MainPage{}, model.ErrStatsError
		default:
			return stats.MainPage{}, model.ErrStatsError
		}
	}
	return stats.MainPage{
		ActiveLeadsAmount:     uint(resp.Data.ActiveLeadsAmount),
		ActiveLeadsPrice:      uint(resp.Data.ActiveLeadsPrice),
		ClosedLeadsAmount:     uint(resp.Data.ClosedLeadsAmount),
		ClosedLeadsPrice:      uint(resp.Data.ClosedLeadsPrice),
		ActiveAdsAmount:       uint(resp.Data.ActiveAdsAmount),
		CompanyAbsoluteRating: resp.Data.CompanyAbsoluteRating,
		CompanyRelativeRating: resp.Data.CompanyRelativeRating,
	}, nil
}
