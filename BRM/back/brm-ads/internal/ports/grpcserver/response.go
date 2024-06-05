package grpcserver

import (
	"brm-ads/internal/model"
	"brm-ads/internal/ports/grpcserver/pb"
	"context"
)

func modelResponseToResponse(resp model.Response) *pb.Response {
	if resp.Id == 0 {
		return nil
	}
	return &pb.Response{
		Id:           resp.Id,
		CompanyId:    resp.CompanyId,
		EmployeeId:   resp.EmployeeId,
		AdId:         resp.AdId,
		CreationDate: resp.CreationDate.UTC().Unix(),
	}
}

func (s *Server) CreateResponse(ctx context.Context, req *pb.CreateResponseRequest) (*pb.CreateResponseResponse, error) {
	resp, err := s.App.CreateResponse(ctx,
		req.CompanyId,
		req.EmployeeId,
		req.AdId,
	)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.CreateResponseResponse{
		Response: modelResponseToResponse(resp),
	}, nil
}

func (s *Server) GetResponses(ctx context.Context, req *pb.GetResponsesRequest) (*pb.GetResponsesResponse, error) {
	responses, amount, err := s.App.GetResponses(ctx,
		req.CompanyId,
		req.EmployeeId,
		uint(req.Limit),
		uint(req.Offset))
	if err != nil {
		return nil, mapErrors(err)
	}

	grpcResp := &pb.GetResponsesResponse{
		List:   make([]*pb.Response, len(responses)),
		Amount: uint64(amount),
	}
	for i, resp := range responses {
		grpcResp.List[i] = modelResponseToResponse(resp)
	}
	return grpcResp, nil
}
