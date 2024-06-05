package grpcserver

import (
	"auth/internal/model"
	"auth/internal/ports/grpcserver/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) RegisterEmployee(ctx context.Context, req *pb.RegisterEmployeeRequest) (*empty.Empty, error) {
	err := s.App.RegisterEmployee(ctx, model.Employee{
		Email:      req.Email,
		Password:   req.Password,
		EmployeeId: req.EmployeeId,
		CompanyId:  req.CompanyId,
	})
	if err != nil {
		return &empty.Empty{}, mapErrors(err)
	}
	return &empty.Empty{}, nil
}

func (s *Server) DeleteEmployee(ctx context.Context, req *pb.DeleteEmployeeRequest) (*empty.Empty, error) {
	err := s.App.DeleteEmployee(ctx, req.Email)
	if err != nil {
		return &empty.Empty{}, mapErrors(err)
	}
	return &empty.Empty{}, nil
}
