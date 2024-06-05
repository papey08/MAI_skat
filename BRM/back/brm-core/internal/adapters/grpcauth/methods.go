package grpcauth

import (
	"brm-core/internal/adapters/grpcauth/pb"
	"brm-core/internal/model"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *authClientImpl) RegisterEmployee(ctx context.Context, creds model.EmployeeCredentials) error {
	_, err := a.cli.RegisterEmployee(ctx, &pb.RegisterEmployeeRequest{
		Email:      creds.Email,
		Password:   creds.Password,
		EmployeeId: creds.EmployeeId,
		CompanyId:  creds.CompanyId,
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.AlreadyExists:
			return model.ErrEmailRegistered
		case codes.FailedPrecondition:
			return model.ErrValidationError
		default:
			return model.ErrAuthServiceError
		}
	}
	return nil
}

func (a *authClientImpl) DeleteEmployee(ctx context.Context, email string) error {
	_, err := a.cli.DeleteEmployee(ctx, &pb.DeleteEmployeeRequest{Email: email})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return model.ErrEmployeeNotExists
		case codes.AlreadyExists:
			return model.ErrEmailRegistered
		case codes.ResourceExhausted:
			return model.ErrAuthServiceError
		default:
			return model.ErrAuthServiceError
		}
	}
	return nil
}
