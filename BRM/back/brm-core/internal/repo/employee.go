package repo

import (
	"brm-core/internal/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	createEmployeeQuery = `
		INSERT INTO "employees" ("company_id", "first_name", "second_name", "email", "job_title", "department", "image_url", "creation_date", "is_deleted") 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING "id";`

	updateEmployeeQuery = `
		UPDATE "employees"
		SET "first_name" = $2,
		    "second_name" = $3,
		    "job_title" = $4,
		    "department" = $5,
		    "image_url" = $6
		WHERE "id" = $1 AND (NOT "is_deleted");`

	deleteEmployeeQuery = `
		UPDATE "employees"
		SET "is_deleted" = true
		WHERE "id" = $1 AND (NOT "is_deleted");`

	getCompanyEmployeesQuery = `
		SELECT * FROM "employees"
		WHERE "company_id" = $1 AND (NOT "is_deleted")
			AND ((NOT $2) OR "job_title" = $3)
			AND ((NOT $4) OR "department" = $5)
		LIMIT $6 OFFSET $7;`

	getCompanyEmployeesAmountQuery = `
		SELECT COUNT(*) FROM "employees"
		WHERE "company_id" = $1 AND (NOT "is_deleted")
			AND ((NOT $2) OR "job_title" = $3)
			AND ((NOT $4) OR "department" = $5);`

	getEmployeeByNameQuery = `
		SELECT * FROM "employees"
		WHERE "company_id" = $1 AND (NOT "is_deleted") AND ("first_name" LIKE $2 OR "second_name" LIKE $2)
		LIMIT $3 OFFSET $4;`

	getEmployeeByNameAmountQuery = `
		SELECT COUNT(*) FROM "employees"
		WHERE "company_id" = $1 AND (NOT "is_deleted") AND ("first_name" LIKE $2 OR "second_name" LIKE $2);`

	getEmployeeByIdQuery = `
		SELECT * FROM "employees"
		WHERE "id" = $1 AND (NOT "is_deleted");`
)

func (c *coreRepoImpl) CreateEmployee(ctx context.Context, employee model.Employee) (model.Employee, error) {
	var employeeId uint64
	var pgErr *pgconn.PgError
	if err := c.QueryRow(ctx, createEmployeeQuery,
		employee.CompanyId,
		employee.FirstName,
		employee.SecondName,
		employee.Email,
		employee.JobTitle,
		employee.Department,
		employee.ImageURL,
		employee.CreationDate,
		employee.IsDeleted,
	).Scan(&employeeId); errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // duplicate primary key error
			return model.Employee{}, model.ErrEmailRegistered
		default:
			return model.Employee{}, model.ErrServiceError
		}
	} else if err != nil {
		return model.Employee{}, errors.Join(model.ErrDatabaseError, err)
	} else {
		employee.Id = employeeId
		return employee, nil
	}
}

func (c *coreRepoImpl) UpdateEmployee(ctx context.Context, employeeId uint64, upd model.UpdateEmployee) (model.Employee, error) {
	if e, err := c.Exec(ctx, updateEmployeeQuery,
		employeeId,
		upd.FirstName,
		upd.SecondName,
		upd.JobTitle,
		upd.Department,
		upd.ImageURL,
	); err != nil {
		return model.Employee{}, errors.Join(model.ErrDatabaseError, err)
	} else if e.RowsAffected() == 0 {
		return model.Employee{}, model.ErrEmployeeNotExists
	} else {
		return c.GetEmployeeById(ctx, employeeId)
	}
}

func (c *coreRepoImpl) DeleteEmployee(ctx context.Context, employeeId uint64) error {
	if e, err := c.Exec(ctx, deleteEmployeeQuery,
		employeeId,
	); err != nil {
		return errors.Join(model.ErrDatabaseError, err)
	} else if e.RowsAffected() == 0 {
		return model.ErrEmployeeNotExists
	} else {
		return nil
	}
}

func (c *coreRepoImpl) GetCompanyEmployees(ctx context.Context, companyId uint64, filter model.FilterEmployee) ([]model.Employee, uint, error) {
	var amount uint
	if err := c.QueryRow(ctx, getCompanyEmployeesAmountQuery,
		companyId,
		filter.ByJobTitle,
		filter.JobTitle,
		filter.ByDepartment,
		filter.Department,
	).Scan(&amount); err != nil {
		return nil, 0, errors.Join(model.ErrDatabaseError, err)
	} else if amount == 0 {
		return []model.Employee{}, amount, nil
	}

	rows, err := c.Query(ctx, getCompanyEmployeesQuery,
		companyId,
		filter.ByJobTitle,
		filter.JobTitle,
		filter.ByDepartment,
		filter.Department,
		filter.Limit,
		filter.Offset)
	if err != nil {
		return []model.Employee{}, 0, errors.Join(model.ErrDatabaseError, err)
	}
	defer rows.Close()

	employees := make([]model.Employee, 0)
	for rows.Next() {
		var e model.Employee
		_ = rows.Scan(
			&e.Id,
			&e.CompanyId,
			&e.FirstName,
			&e.SecondName,
			&e.Email,
			&e.JobTitle,
			&e.Department,
			&e.ImageURL,
			&e.CreationDate,
			&e.IsDeleted)
		employees = append(employees, e)
	}
	return employees, amount, nil
}

func (c *coreRepoImpl) GetEmployeeByName(ctx context.Context, companyId uint64, ebn model.EmployeeByName) ([]model.Employee, uint, error) {
	var amount uint
	if err := c.QueryRow(ctx, getEmployeeByNameAmountQuery,
		companyId,
		ebn.Pattern+"%",
	).Scan(&amount); err != nil {
		return nil, 0, errors.Join(model.ErrDatabaseError, err)
	} else if amount == 0 {
		return []model.Employee{}, amount, nil
	}

	rows, err := c.Query(ctx, getEmployeeByNameQuery,
		companyId,
		ebn.Pattern+"%",
		ebn.Limit,
		ebn.Offset)
	if err != nil {
		return []model.Employee{}, 0, errors.Join(model.ErrDatabaseError, err)
	}
	defer rows.Close()

	employees := make([]model.Employee, 0)
	for rows.Next() {
		var e model.Employee
		_ = rows.Scan(
			&e.Id,
			&e.CompanyId,
			&e.FirstName,
			&e.SecondName,
			&e.Email,
			&e.JobTitle,
			&e.Department,
			&e.ImageURL,
			&e.CreationDate,
			&e.IsDeleted)
		employees = append(employees, e)
	}
	return employees, amount, nil
}

func (c *coreRepoImpl) GetEmployeeById(ctx context.Context, employeeId uint64) (model.Employee, error) {
	row := c.QueryRow(ctx, getEmployeeByIdQuery, employeeId)
	var employee model.Employee
	if err := row.Scan(
		&employee.Id,
		&employee.CompanyId,
		&employee.FirstName,
		&employee.SecondName,
		&employee.Email,
		&employee.JobTitle,
		&employee.Department,
		&employee.ImageURL,
		&employee.CreationDate,
		&employee.IsDeleted,
	); errors.Is(err, pgx.ErrNoRows) {
		return model.Employee{}, model.ErrEmployeeNotExists
	} else if err != nil {
		return model.Employee{}, errors.Join(model.ErrDatabaseError, err)
	} else {
		return employee, nil
	}
}
