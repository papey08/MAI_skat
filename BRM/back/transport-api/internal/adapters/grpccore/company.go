package grpccore

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"transport-api/internal/adapters/grpccore/pb"
	"transport-api/internal/model"
	"transport-api/internal/model/core"
)

func respToCompany(company *pb.Company) core.Company {
	if company == nil {
		return core.Company{}
	}
	return core.Company{
		Id:           company.Id,
		Name:         company.Name,
		Description:  company.Description,
		Industry:     company.Industry,
		OwnerId:      company.OwnerId,
		Rating:       company.Rating,
		CreationDate: company.CreationDate,
		IsDeleted:    company.IsDeleted,
	}
}

func (c *coreClientImpl) GetCompany(ctx context.Context, id uint64) (core.Company, error) {
	resp, err := c.cli.GetCompany(ctx, &pb.GetCompanyRequest{Id: id})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return core.Company{}, model.ErrCompanyNotExists
		case codes.ResourceExhausted:
			return core.Company{}, model.ErrCoreError
		default:
			return core.Company{}, model.ErrCoreUnknown
		}
	}
	return respToCompany(resp.Company), nil
}

func (c *coreClientImpl) UpdateCompany(ctx context.Context, companyId uint64, ownerId uint64, upd core.UpdateCompany) (core.Company, error) {
	resp, err := c.cli.UpdateCompany(ctx, &pb.UpdateCompanyRequest{
		CompanyId: companyId,
		OwnerId:   ownerId,
		Upd: &pb.UpdateCompanyFields{
			Name:        upd.Name,
			Description: upd.Description,
			Industry:    upd.Industry,
			OwnerId:     upd.OwnerId,
		},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			// костыль, ну а чё поделать
			switch {
			case strings.Contains(err.Error(), "company"):
				return core.Company{}, model.ErrCompanyNotExists
			case strings.Contains(err.Error(), "employee"):
				return core.Company{}, model.ErrEmployeeNotExists
			case strings.Contains(err.Error(), "industry"):
				return core.Company{}, model.ErrIndustryNotExists
			}
		case codes.PermissionDenied:
			return core.Company{}, model.ErrPermissionDenied
		case codes.FailedPrecondition:
			return core.Company{}, model.ErrInvalidInput
		case codes.ResourceExhausted:
			return core.Company{}, model.ErrCoreError
		default:
			return core.Company{}, model.ErrCoreUnknown
		}
	}
	return respToCompany(resp.Company), nil
}

func (c *coreClientImpl) GetIndustries(ctx context.Context) (map[string]uint64, error) {
	resp, err := c.cli.GetIndustries(ctx, &empty.Empty{})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.ResourceExhausted:
			return map[string]uint64{}, model.ErrCoreError
		default:
			return map[string]uint64{}, model.ErrCoreUnknown
		}
	}
	return resp.Data, nil
}
