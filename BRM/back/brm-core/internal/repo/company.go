package repo

import (
	"brm-core/internal/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	getCompanyQuery = `
		SELECT "companies"."id", "companies"."name", "description", "industries"."name", "owner_id", "rating", "creation_date", "is_deleted" FROM "companies"
		INNER JOIN "industries" ON "companies"."industry" = "industries"."id"
		WHERE "companies"."id" = $1 AND (NOT "is_deleted");`

	createCompanyQuery = `
		INSERT INTO "companies" ("name", "description", "industry", "owner_id", "rating", "creation_date", "is_deleted") 
		VALUES ($1, $2, (SELECT "id" FROM "industries" WHERE "name" = $3), $4, $5, $6, $7)
		RETURNING "id";`

	createOwnerQuery = `
		INSERT INTO "employees" ("company_id", "first_name", "second_name", "email", "job_title", "department", "creation_date", "is_deleted")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING "id";`

	setCompanyOwnerIdQuery = `
		UPDATE "companies"
		SET "owner_id" = $2
		WHERE "id" = $1;`

	updateCompanyQuery = `
		UPDATE "companies"
		SET "name" = $2,
		    "description" = $3,
		    "industry" = (SELECT "id" FROM "industries" WHERE "name" = $4),
		    "owner_id" = $5
		WHERE "id" = $1 AND (NOT "is_deleted");`

	getIndustriesQuery = `
		SELECT * FROM "industries";`
)

func (c *coreRepoImpl) GetCompany(ctx context.Context, id uint64) (model.Company, error) {
	row := c.QueryRow(ctx, getCompanyQuery, id)
	var company model.Company
	if err := row.Scan(
		&company.Id,
		&company.Name,
		&company.Description,
		&company.Industry,
		&company.OwnerId,
		&company.Rating,
		&company.CreationDate,
		&company.IsDeleted,
	); errors.Is(err, pgx.ErrNoRows) {
		return model.Company{}, model.ErrCompanyNotExists
	} else if err != nil {
		return model.Company{}, errors.Join(model.ErrDatabaseError, err)
	} else {
		return company, nil
	}
}

func (c *coreRepoImpl) CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error) {
	var companyId, ownerId uint64
	var pgErr *pgconn.PgError
	if err := c.QueryRow(ctx, createCompanyQuery,
		company.Name,
		company.Description,
		company.Industry,
		0,
		company.Rating,
		company.CreationDate,
		company.IsDeleted,
	).Scan(&companyId); errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23502":
			return model.Company{}, model.Employee{}, model.ErrIndustryNotExists
		default:
			return model.Company{}, model.Employee{}, errors.Join(model.ErrDatabaseError, err)
		}
	}

	company.Id = companyId
	owner.CompanyId = companyId

	if err := c.QueryRow(ctx, createOwnerQuery,
		owner.CompanyId,
		owner.FirstName,
		owner.SecondName,
		owner.Email,
		owner.JobTitle,
		owner.Department,
		owner.CreationDate,
		owner.IsDeleted,
	).Scan(&ownerId); errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // duplicate primary key error
			return model.Company{}, model.Employee{}, model.ErrEmailRegistered
		default:
			return model.Company{}, model.Employee{}, model.ErrServiceError
		}
	} else if err != nil {
		return model.Company{}, model.Employee{}, errors.Join(model.ErrDatabaseError, err)
	}

	owner.Id = ownerId
	company.OwnerId = ownerId

	if _, err := c.Exec(ctx, setCompanyOwnerIdQuery, companyId, ownerId); err != nil {
		return model.Company{}, model.Employee{}, model.ErrDatabaseError
	}

	return company, owner, nil
}

func (c *coreRepoImpl) UpdateCompany(ctx context.Context, companyId uint64, upd model.UpdateCompany) (model.Company, error) {
	var pgErr *pgconn.PgError
	if e, err := c.Exec(ctx, updateCompanyQuery,
		companyId,
		upd.Name,
		upd.Description,
		upd.Industry,
		upd.OwnerId,
	); errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23502": // null column industry
			return model.Company{}, model.ErrIndustryNotExists
		default:
			return model.Company{}, errors.Join(model.ErrDatabaseError, err)
		}
	} else if e.RowsAffected() == 0 {
		return model.Company{}, model.ErrCompanyNotExists
	} else {
		return c.GetCompany(ctx, companyId)
	}
}

func (c *coreRepoImpl) GetIndustries(ctx context.Context) (map[string]uint64, error) {
	rows, err := c.Query(ctx, getIndustriesQuery)
	if err != nil {
		return map[string]uint64{}, model.ErrDatabaseError
	}
	defer rows.Close()

	industries := make(map[string]uint64)
	for rows.Next() {
		var id uint64
		var industry string
		_ = rows.Scan(&id, &industry)
		industries[industry] = id
	}
	return industries, nil
}
