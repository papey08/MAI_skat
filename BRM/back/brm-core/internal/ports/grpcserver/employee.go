package grpcserver

import (
	"brm-core/internal/model"
	"brm-core/internal/ports/grpcserver/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"time"
)

func employeeToModelEmployee(employee *pb.Employee) model.Employee {
	if employee == nil {
		return model.Employee{}
	}
	return model.Employee{
		Id:           employee.Id,
		CompanyId:    employee.CompanyId,
		FirstName:    employee.FirstName,
		SecondName:   employee.SecondName,
		Email:        employee.Email,
		Password:     employee.Password,
		JobTitle:     employee.JobTitle,
		Department:   employee.Department,
		ImageURL:     employee.ImageUrl,
		CreationDate: time.Unix(employee.CreationDate, 0),
		IsDeleted:    employee.IsDeleted,
	}
}

func modelEmployeeToEmployee(employee model.Employee) *pb.Employee {
	if employee.Id == 0 {
		return nil
	}
	return &pb.Employee{
		Id:           employee.Id,
		CompanyId:    employee.CompanyId,
		FirstName:    employee.FirstName,
		SecondName:   employee.SecondName,
		Email:        employee.Email,
		Password:     employee.Password,
		JobTitle:     employee.JobTitle,
		Department:   employee.Department,
		ImageUrl:     employee.ImageURL,
		CreationDate: employee.CreationDate.UTC().Unix(),
		IsDeleted:    employee.IsDeleted,
	}
}

func (s *Server) CreateEmployee(ctx context.Context, req *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {
	employee, err := s.App.CreateEmployee(ctx,
		req.CompanyId,
		req.OwnerId,
		employeeToModelEmployee(req.Employee),
	)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.CreateEmployeeResponse{
		Employee: modelEmployeeToEmployee(employee),
	}, nil
}

func (s *Server) UpdateEmployee(ctx context.Context, req *pb.UpdateEmployeeRequest) (*pb.UpdateEmployeeResponse, error) {
	employee, err := s.App.UpdateEmployee(ctx,
		req.CompanyId,
		req.OwnerId,
		req.EmployeeId,
		model.UpdateEmployee{
			FirstName:  req.Upd.FirstName,
			SecondName: req.Upd.SecondName,
			JobTitle:   req.Upd.JobTitle,
			Department: req.Upd.Department,
			ImageURL:   req.Upd.ImageUrl,
		},
	)
	if err != nil {
		return nil, mapErrors(err)
	}
	return &pb.UpdateEmployeeResponse{
		Employee: modelEmployeeToEmployee(employee),
	}, nil
}

func (s *Server) DeleteEmployee(ctx context.Context, req *pb.DeleteEmployeeRequest) (*empty.Empty, error) {
	if err := s.App.DeleteEmployee(ctx,
		req.CompanyId,
		req.OwnerId,
		req.EmployeeId,
	); err != nil {
		return nil, mapErrors(err)
	}
	return &empty.Empty{}, nil
}

func (s *Server) GetCompanyEmployees(ctx context.Context, req *pb.GetCompanyEmployeesRequest) (*pb.GetCompanyEmployeesResponse, error) {
	employees, amount, err := s.App.GetCompanyEmployees(
		ctx,
		req.CompanyId,
		req.EmployeeId,
		model.FilterEmployee{
			ByJobTitle:   req.Filter.ByJobTitle,
			JobTitle:     req.Filter.JobTitle,
			ByDepartment: req.Filter.ByDepartment,
			Department:   req.Filter.Department,
			Limit:        uint(req.Filter.Limit),
			Offset:       uint(req.Filter.Offset),
		})
	if err != nil {
		return nil, mapErrors(err)
	}

	resp := &pb.GetCompanyEmployeesResponse{
		List:   make([]*pb.Employee, len(employees)),
		Amount: uint64(amount),
	}
	for i, empl := range employees {
		resp.List[i] = modelEmployeeToEmployee(empl)
	}
	return resp, nil
}

func (s *Server) GetEmployeeByName(ctx context.Context, req *pb.GetEmployeeByNameRequest) (*pb.GetEmployeeByNameResponse, error) {
	employees, amount, err := s.App.GetEmployeeByName(
		ctx,
		req.CompanyId,
		req.EmployeeId,
		model.EmployeeByName{
			Pattern: req.Ebn.Pattern,
			Limit:   uint(req.Ebn.Limit),
			Offset:  uint(req.Ebn.Offset),
		})
	if err != nil {
		return nil, mapErrors(err)
	}

	resp := &pb.GetEmployeeByNameResponse{
		List:   make([]*pb.Employee, len(employees)),
		Amount: uint64(amount),
	}
	for i, empl := range employees {
		resp.List[i] = modelEmployeeToEmployee(empl)
	}
	return resp, nil
}

func (s *Server) GetEmployeeById(ctx context.Context, req *pb.GetEmployeeByIdRequest) (*pb.GetEmployeeByIdResponse, error) {
	employee, err := s.App.GetEmployeeById(
		ctx,
		req.CompanyId,
		req.EmployeeId,
		req.EmployeeIdToFind,
	)
	if err != nil {
		return nil, mapErrors(err)
	}

	return &pb.GetEmployeeByIdResponse{
		Employee: modelEmployeeToEmployee(employee),
	}, nil
}
