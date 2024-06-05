package grpccore

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"registration/internal/adapters/grpccore/pb"
	"registration/internal/model"
)

type coreClientImpl struct {
	cli pb.CoreServiceClient
}

func NewCoreClient(ctx context.Context, addr string) (CoreClient, error) {
	if conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return &coreClientImpl{}, fmt.Errorf("grpc core client: %w", err)
	} else {
		return &coreClientImpl{
			cli: pb.NewCoreServiceClient(conn),
		}, nil
	}
}

func (c *coreClientImpl) CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error) {
	resp, err := c.cli.CreateCompanyAndOwner(ctx, &pb.CreateCompanyAndOwnerRequest{
		Company: &pb.Company{
			Id:           company.Id,
			Name:         company.Name,
			Description:  company.Description,
			Industry:     company.Industry,
			OwnerId:      company.OwnerId,
			Rating:       company.Rating,
			CreationDate: company.CreationDate,
			IsDeleted:    company.IsDeleted,
		},
		Owner: &pb.Employee{
			Id:           owner.Id,
			CompanyId:    owner.CompanyId,
			FirstName:    owner.FirstName,
			SecondName:   owner.SecondName,
			Email:        owner.Email,
			Password:     owner.Password,
			JobTitle:     owner.JobTitle,
			Department:   owner.Department,
			CreationDate: owner.CreationDate,
			IsDeleted:    owner.IsDeleted,
		},
	})

	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return model.Company{}, model.Employee{}, model.ErrIndustryNotExists
		case codes.ResourceExhausted:
			return model.Company{}, model.Employee{}, model.ErrCoreError
		case codes.FailedPrecondition:
			return model.Company{}, model.Employee{}, model.ErrInvalidInput
		default:
			return model.Company{}, model.Employee{}, model.ErrCoreUnknown
		}
	}

	return model.Company{
			Id:           resp.Company.Id,
			Name:         resp.Company.Name,
			Description:  resp.Company.Description,
			Industry:     resp.Company.Industry,
			OwnerId:      resp.Company.OwnerId,
			Rating:       resp.Company.Rating,
			CreationDate: resp.Company.CreationDate,
			IsDeleted:    resp.Company.IsDeleted,
		},
		model.Employee{
			Id:           resp.Owner.Id,
			CompanyId:    resp.Owner.CompanyId,
			FirstName:    resp.Owner.FirstName,
			SecondName:   resp.Owner.SecondName,
			Email:        resp.Owner.Email,
			Password:     resp.Owner.Password,
			JobTitle:     resp.Owner.JobTitle,
			Department:   resp.Owner.Department,
			CreationDate: resp.Owner.CreationDate,
			IsDeleted:    resp.Owner.IsDeleted,
		}, nil
}

func (c *coreClientImpl) GetIndustriesList(ctx context.Context) (map[string]uint64, error) {
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
