package grpcads

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"transport-api/internal/adapters/grpcads/pb"
	"transport-api/internal/model"
	"transport-api/internal/model/ads"
)

func respToResponse(resp *pb.Response) ads.Response {
	if resp == nil {
		return ads.Response{}
	}
	return ads.Response{
		Id:           resp.Id,
		CompanyId:    resp.CompanyId,
		EmployeeId:   resp.EmployeeId,
		AdId:         resp.AdId,
		CreationDate: resp.CreationDate,
	}
}

func (a *adsClientImpl) CreateResponse(ctx context.Context, companyId uint64, employeeId uint64, adId uint64) (ads.Response, error) {
	resp, err := a.cli.CreateResponse(ctx, &pb.CreateResponseRequest{
		CompanyId:  companyId,
		EmployeeId: employeeId,
		AdId:       adId,
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return ads.Response{}, model.ErrAdNotExists
		case codes.FailedPrecondition:
			return ads.Response{}, model.ErrSameCompany
		case codes.ResourceExhausted:
			return ads.Response{}, model.ErrAdsError
		case codes.Unknown:
			return ads.Response{}, model.ErrAdsUnknown
		}
	}
	return respToResponse(resp.Response), nil
}

func (a *adsClientImpl) GetResponses(ctx context.Context, companyId uint64, employeeId uint64, limit uint, offset uint) ([]ads.Response, uint, error) {
	grpcResp, err := a.cli.GetResponses(ctx, &pb.GetResponsesRequest{
		CompanyId:  companyId,
		EmployeeId: employeeId,
		Limit:      uint64(limit),
		Offset:     uint64(offset),
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.ResourceExhausted:
			return []ads.Response{}, 0, model.ErrAdsError
		case codes.Unknown:
			return []ads.Response{}, 0, model.ErrAdsUnknown
		}
	}
	responses := make([]ads.Response, len(grpcResp.List))
	for i, resp := range grpcResp.List {
		responses[i] = respToResponse(resp)
	}
	return responses, uint(grpcResp.Amount), nil
}
