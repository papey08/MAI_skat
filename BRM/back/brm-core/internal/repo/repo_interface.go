package repo

import (
	"brm-core/internal/model"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CoreRepo interface {
	CompanyRepo
	EmployeeRepo
	ContactRepo
}

func New(pool *pgxpool.Pool) CoreRepo {
	return &coreRepoImpl{
		Pool: pool,
	}
}

type CompanyRepo interface {
	GetCompany(ctx context.Context, id uint64) (model.Company, error)
	CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error)
	UpdateCompany(ctx context.Context, companyId uint64, upd model.UpdateCompany) (model.Company, error)

	GetIndustries(ctx context.Context) (map[string]uint64, error)
}

type EmployeeRepo interface {
	CreateEmployee(ctx context.Context, employee model.Employee) (model.Employee, error)
	UpdateEmployee(ctx context.Context, employeeId uint64, upd model.UpdateEmployee) (model.Employee, error)
	DeleteEmployee(ctx context.Context, employeeId uint64) error

	GetCompanyEmployees(ctx context.Context, companyId uint64, filter model.FilterEmployee) ([]model.Employee, uint, error)
	GetEmployeeByName(ctx context.Context, companyId uint64, ebn model.EmployeeByName) ([]model.Employee, uint, error)
	GetEmployeeById(ctx context.Context, employeeId uint64) (model.Employee, error)
}

type ContactRepo interface {
	CreateContact(ctx context.Context, contact model.Contact) (model.Contact, error)
	UpdateContact(ctx context.Context, ownerId uint64, contactId uint64, upd model.UpdateContact) (model.Contact, error)
	DeleteContact(ctx context.Context, ownerId uint64, contactId uint64) error

	GetContacts(ctx context.Context, ownerId uint64, pagination model.GetContacts) ([]model.Contact, error)
	GetContactsAmount(ctx context.Context, ownerId uint64) (uint, error)
	GetContactById(ctx context.Context, ownerId uint64, contactId uint64) (model.Contact, error)
}
