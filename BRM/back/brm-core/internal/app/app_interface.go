package app

import (
	"brm-core/internal/adapters/grpcauth"
	"brm-core/internal/model"
	"brm-core/internal/repo"
	"brm-core/pkg/logger"
	"context"
)

type App interface {
	CompanyApp
	EmployeeApp
	ContactApp
}

func New(coreRepo repo.CoreRepo, authCli grpcauth.AuthClient, logs logger.Logger) App {
	return &appImpl{
		coreRepo: coreRepo,
		auth:     authCli,
		logs:     logs,
	}
}

type CompanyApp interface {
	GetCompany(ctx context.Context, id uint64) (model.Company, error)
	CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error)
	UpdateCompany(ctx context.Context, companyId uint64, ownerId uint64, upd model.UpdateCompany) (model.Company, error)

	GetIndustries(ctx context.Context) (map[string]uint64, error)
}

type EmployeeApp interface {
	CreateEmployee(ctx context.Context, companyId uint64, ownerId uint64, employee model.Employee) (model.Employee, error)
	UpdateEmployee(ctx context.Context, companyId uint64, ownerId uint64, employeeId uint64, upd model.UpdateEmployee) (model.Employee, error)
	DeleteEmployee(ctx context.Context, companyId uint64, ownerId uint64, employeeId uint64) error

	GetCompanyEmployees(ctx context.Context, companyId uint64, employeeId uint64, filter model.FilterEmployee) ([]model.Employee, uint, error)
	GetEmployeeByName(ctx context.Context, companyId uint64, employeeId uint64, ebn model.EmployeeByName) ([]model.Employee, uint, error)
	GetEmployeeById(ctx context.Context, companyId uint64, employeeId uint64, employeeIdToFind uint64) (model.Employee, error)
}

type ContactApp interface {
	CreateContact(ctx context.Context, ownerId uint64, employeeId uint64) (model.Contact, error)
	UpdateContact(ctx context.Context, ownerId uint64, contactId uint64, upd model.UpdateContact) (model.Contact, error)
	DeleteContact(ctx context.Context, ownerId uint64, contactId uint64) error

	GetContacts(ctx context.Context, ownerId uint64, pagination model.GetContacts) ([]model.Contact, uint, error)
	GetContactById(ctx context.Context, ownerId uint64, contactId uint64) (model.Contact, error)
}
