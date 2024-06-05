package grpccore

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"transport-api/internal/adapters/grpccore/pb"
	"transport-api/internal/model"
	"transport-api/internal/model/core"
)

func respToEmployee(employee *pb.Employee) core.Employee {
	if employee == nil {
		return core.Employee{}
	}
	return core.Employee{
		Id:           employee.Id,
		CompanyId:    employee.CompanyId,
		FirstName:    employee.FirstName,
		SecondName:   employee.SecondName,
		Email:        employee.Email,
		Password:     employee.Password,
		JobTitle:     employee.JobTitle,
		Department:   employee.Department,
		ImageURL:     employee.ImageUrl,
		CreationDate: employee.CreationDate,
		IsDeleted:    employee.IsDeleted,
	}
}

func employeeToRequest(employee core.Employee) *pb.Employee {
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
		CreationDate: employee.CreationDate,
		IsDeleted:    employee.IsDeleted,
	}
}

func (c *coreClientImpl) CreateEmployee(ctx context.Context, companyId uint64, ownerId uint64, employee core.Employee) (core.Employee, error) {
	resp, err := c.cli.CreateEmployee(ctx, &pb.CreateEmployeeRequest{
		CompanyId: companyId,
		OwnerId:   ownerId,
		Employee:  employeeToRequest(employee),
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.PermissionDenied:
			return core.Employee{}, model.ErrPermissionDenied
		case codes.AlreadyExists:
			return core.Employee{}, model.ErrEmailRegistered
		case codes.FailedPrecondition:
			return core.Employee{}, model.ErrInvalidInput
		case codes.ResourceExhausted:
			return core.Employee{}, model.ErrCoreError
		default:
			return core.Employee{}, model.ErrCoreUnknown
		}
	}
	return respToEmployee(resp.Employee), nil
}

func (c *coreClientImpl) UpdateEmployee(ctx context.Context, companyId uint64, ownerId uint64, employeeId uint64, upd core.UpdateEmployee) (core.Employee, error) {
	resp, err := c.cli.UpdateEmployee(ctx, &pb.UpdateEmployeeRequest{
		CompanyId:  companyId,
		OwnerId:    ownerId,
		EmployeeId: employeeId,
		Upd: &pb.UpdateEmployeeFields{
			FirstName:  upd.FirstName,
			SecondName: upd.SecondName,
			JobTitle:   upd.JobTitle,
			Department: upd.Department,
			ImageUrl:   upd.ImageURL,
		},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return core.Employee{}, model.ErrEmployeeNotExists
		case codes.PermissionDenied:
			return core.Employee{}, model.ErrPermissionDenied
		case codes.FailedPrecondition:
			return core.Employee{}, model.ErrInvalidInput
		case codes.ResourceExhausted:
			return core.Employee{}, model.ErrCoreError
		default:
			return core.Employee{}, model.ErrCoreUnknown
		}
	}
	return respToEmployee(resp.Employee), nil
}

func (c *coreClientImpl) DeleteEmployee(ctx context.Context, companyId uint64, ownerId uint64, employeeId uint64) error {
	_, err := c.cli.DeleteEmployee(ctx, &pb.DeleteEmployeeRequest{
		CompanyId:  companyId,
		OwnerId:    ownerId,
		EmployeeId: employeeId,
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return model.ErrEmployeeNotExists
		case codes.PermissionDenied:
			return model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return model.ErrCoreError
		case codes.Aborted:
			return model.ErrOwnerDeletion
		default:
			return model.ErrCoreUnknown
		}
	}
	return nil
}

func (c *coreClientImpl) GetCompanyEmployees(ctx context.Context, companyId uint64, employeeId uint64, filter core.FilterEmployee) ([]core.Employee, uint, error) {
	resp, err := c.cli.GetCompanyEmployees(ctx, &pb.GetCompanyEmployeesRequest{
		CompanyId:  companyId,
		EmployeeId: employeeId,
		Filter: &pb.FilterEmployee{
			ByJobTitle:   filter.ByJobTitle,
			JobTitle:     filter.JobTitle,
			ByDepartment: filter.ByDepartment,
			Department:   filter.Department,
			Limit:        int64(filter.Limit),
			Offset:       int64(filter.Offset),
		},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return []core.Employee{}, 0, model.ErrEmployeeNotExists
		case codes.PermissionDenied:
			return []core.Employee{}, 0, model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return []core.Employee{}, 0, model.ErrCoreError
		default:
			return []core.Employee{}, 0, model.ErrCoreUnknown
		}
	}
	employees := make([]core.Employee, len(resp.List))
	for i, empl := range resp.List {
		employees[i] = respToEmployee(empl)
	}
	return employees, uint(resp.Amount), nil
}

func (c *coreClientImpl) GetEmployeeByName(ctx context.Context, companyId uint64, employeeId uint64, ebn core.EmployeeByName) ([]core.Employee, uint, error) {
	resp, err := c.cli.GetEmployeeByName(ctx, &pb.GetEmployeeByNameRequest{
		CompanyId:  companyId,
		EmployeeId: employeeId,
		Ebn: &pb.EmployeeByName{
			Pattern: ebn.Pattern,
			Limit:   int64(ebn.Limit),
			Offset:  int64(ebn.Offset),
		},
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return []core.Employee{}, 0, model.ErrEmployeeNotExists
		case codes.PermissionDenied:
			return []core.Employee{}, 0, model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return []core.Employee{}, 0, model.ErrCoreError
		default:
			return []core.Employee{}, 0, model.ErrCoreUnknown
		}
	}
	employees := make([]core.Employee, len(resp.List))
	for i, empl := range resp.List {
		employees[i] = respToEmployee(empl)
	}
	return employees, uint(resp.Amount), nil
}

func (c *coreClientImpl) GetEmployeeById(ctx context.Context, companyId uint64, employeeId uint64, employeeIdToFind uint64) (core.Employee, error) {
	resp, err := c.cli.GetEmployeeById(ctx, &pb.GetEmployeeByIdRequest{
		CompanyId:        companyId,
		EmployeeId:       employeeId,
		EmployeeIdToFind: employeeIdToFind,
	})
	if err != nil {
		code := status.Code(err)
		switch code {
		case codes.NotFound:
			return core.Employee{}, model.ErrEmployeeNotExists
		case codes.PermissionDenied:
			return core.Employee{}, model.ErrPermissionDenied
		case codes.ResourceExhausted:
			return core.Employee{}, model.ErrCoreError
		default:
			return core.Employee{}, model.ErrCoreUnknown
		}
	}
	empl := respToEmployee(resp.Employee)
	return empl, nil
}
