package grpcserver

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"stats/internal/ports/grpcserver/pb"
)

func (s *Server) GetCompanyMainPage(ctx context.Context, req *pb.GetCompanyMainPageRequest) (*pb.GetCompanyMainPageResponse, error) {
	resp, err := s.a.GetCompanyMainPageStats(ctx, req.CompanyId)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.GetCompanyMainPageResponse{Data: &pb.Data{
		ActiveLeadsAmount:     uint64(resp.ActiveLeadsAmount),
		ActiveLeadsPrice:      uint64(resp.ActiveLeadsPrice),
		ClosedLeadsAmount:     uint64(resp.ClosedLeadsAmount),
		ClosedLeadsPrice:      uint64(resp.ClosedLeadsPrice),
		ActiveAdsAmount:       uint64(resp.ActiveAdsAmount),
		CompanyAbsoluteRating: resp.CompanyAbsoluteRating,
		CompanyRelativeRating: resp.CompanyRelativeRating,
	}}, nil
}

func (s *Server) SubmitClosedLead(ctx context.Context, req *pb.SubmitClosedLeadRequest) (*empty.Empty, error) {
	err := s.a.UpdateRatingByClosedLead(ctx, req.CompanyId, req.Submit)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &empty.Empty{}, nil
}
