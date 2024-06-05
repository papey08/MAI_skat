package grpccore

import (
	"brm-ads/internal/adapters/grpccore/pb"
	"brm-ads/internal/model"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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

func (c *coreClientImpl) GetCompany(ctx context.Context, id uint64) (uint64, error) {
	resp, err := c.cli.GetCompany(ctx, &pb.GetCompanyRequest{Id: id})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return 0, model.ErrCompanyNotExists
		default:
			return 0, model.ErrCoreError
		}
	}
	return resp.Company.Id, nil
}

func (c *coreClientImpl) GetEmployeeById(ctx context.Context, companyId uint64, employeeId uint64, employeeIdToFind uint64) (uint64, uint64, error) {
	resp, err := c.cli.GetEmployeeById(ctx, &pb.GetEmployeeByIdRequest{
		CompanyId:        companyId,
		EmployeeId:       employeeId,
		EmployeeIdToFind: employeeIdToFind,
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return 0, 0, model.ErrEmployeeNotExists
		case codes.PermissionDenied:
			return 0, 0, model.ErrAuthorization
		default:
			return 0, 0, model.ErrCoreError
		}
	}
	return resp.Employee.CompanyId, resp.Employee.Id, nil
}
