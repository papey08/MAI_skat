package repo

import (
	"brm-leads/internal/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	createLeadQuery = `
		INSERT INTO "leads" ("ad_id", "title", "description", "price", "status", "responsible", "company_id", "client_company", "client_employee", "creation_date", "is_deleted")
		VALUES ($1, $2, $3, $4, (SELECT "id" FROM "statuses" WHERE "name" = $5), $6, $7, $8, $9, $10, $11)
		RETURNING "id";`

	getLeadsQuery = `
		SELECT  "leads"."id", 
		        "ad_id", 
		        "title", 
		        "description", 
		        "price", 
		        "statuses"."name", 
		        "responsible", 
		        "company_id",
		        "client_company",
		        "client_employee",
		        "creation_date",
		        "is_deleted" 
		FROM "leads"
		INNER JOIN "statuses" ON "leads"."status" = "statuses"."id"                      
		WHERE "company_id" = $1 AND (NOT "is_deleted")
			AND ((NOT $2) OR "status" = (SELECT "id" FROM "statuses" WHERE "name" = $3))
			AND ((NOT $4) OR "responsible" = $5)
		ORDER BY "creation_date" DESC
		LIMIT $6 OFFSET $7;`

	getLeadsAmountQuery = `
		SELECT  COUNT(*) FROM "leads"
		INNER JOIN "statuses" ON "leads"."status" = "statuses"."id"                      
		WHERE "company_id" = $1 AND (NOT "is_deleted")
			AND ((NOT $2) OR "status" = (SELECT "id" FROM "statuses" WHERE "name" = $3))
			AND ((NOT $4) OR "responsible" = $5);`

	getLeadByIdQuery = `
		SELECT  "leads"."id", 
		        "ad_id", 
		        "title", 
		        "description", 
		        "price", 
		        "statuses"."name", 
		        "responsible", 
		        "company_id",
		        "client_company",
		        "client_employee",
		        "creation_date",
		        "is_deleted" 
		FROM "leads"
		INNER JOIN "statuses" ON "leads"."status" = "statuses"."id" 
		WHERE "leads"."id" = $1 AND (NOT "is_deleted");`

	updateLeadQuery = `
		UPDATE "leads"
		SET "title" = $2,
			"description" = $3,
			"price" = $4,
			"status" = (SELECT "id" FROM "statuses" WHERE "name" = $5),
			"responsible" = $6
		WHERE "id" = $1 AND (NOT "is_deleted");`
)

func (l *leadRepoImpl) CreateLead(ctx context.Context, lead model.Lead) (model.Lead, error) {
	var leadId uint64
	var pgErr *pgconn.PgError
	if err := l.QueryRow(ctx, createLeadQuery,
		lead.AdId,
		lead.Title,
		lead.Description,
		lead.Price,
		lead.Status,
		lead.Responsible,
		lead.CompanyId,
		lead.ClientCompany,
		lead.ClientEmployee,
		lead.CreationDate,
		lead.IsDeleted,
	).Scan(&leadId); errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23502": // null column status
			return model.Lead{}, model.ErrStatusNotExists
		default:
			return model.Lead{}, errors.Join(model.ErrDatabaseError, err)
		}
	} else {
		lead.Id = leadId
		return lead, nil
	}
}

func (l *leadRepoImpl) GetLeads(ctx context.Context, companyId uint64, filter model.Filter) ([]model.Lead, uint, error) {
	var amount uint
	if err := l.QueryRow(ctx, getLeadsAmountQuery,
		companyId,
		filter.ByStatus,
		filter.Status,
		filter.ByResponsible,
		filter.Responsible,
	).Scan(&amount); err != nil {
		return []model.Lead{}, 0, errors.Join(model.ErrDatabaseError, err)
	} else if amount == 0 {
		return []model.Lead{}, 0, nil
	}

	rows, err := l.Query(ctx, getLeadsQuery,
		companyId,
		filter.ByStatus,
		filter.Status,
		filter.ByResponsible,
		filter.Responsible,
		filter.Limit,
		filter.Offset,
	)
	if err != nil {
		return []model.Lead{}, 0, errors.Join(model.ErrDatabaseError, err)
	}
	defer rows.Close()

	leads := make([]model.Lead, 0)
	for rows.Next() {
		var lead model.Lead
		_ = rows.Scan(
			&lead.Id,
			&lead.AdId,
			&lead.Title,
			&lead.Description,
			&lead.Price,
			&lead.Status,
			&lead.Responsible,
			&lead.CompanyId,
			&lead.ClientCompany,
			&lead.ClientEmployee,
			&lead.CreationDate,
			&lead.IsDeleted,
		)
		leads = append(leads, lead)
	}
	return leads, amount, nil
}

func (l *leadRepoImpl) GetLeadById(ctx context.Context, id uint64) (model.Lead, error) {
	row := l.QueryRow(ctx, getLeadByIdQuery, id)
	var lead model.Lead
	if err := row.Scan(
		&lead.Id,
		&lead.AdId,
		&lead.Title,
		&lead.Description,
		&lead.Price,
		&lead.Status,
		&lead.Responsible,
		&lead.CompanyId,
		&lead.ClientCompany,
		&lead.ClientEmployee,
		&lead.CreationDate,
		&lead.IsDeleted,
	); errors.Is(err, pgx.ErrNoRows) {
		return model.Lead{}, model.ErrLeadNotExists
	} else if err != nil {
		return model.Lead{}, errors.Join(model.ErrDatabaseError, err)
	} else {
		return lead, nil
	}
}

func (l *leadRepoImpl) UpdateLead(ctx context.Context, id uint64, upd model.UpdateLead) (model.Lead, error) {
	var pgErr *pgconn.PgError
	if e, err := l.Exec(ctx, updateLeadQuery,
		id,
		upd.Title,
		upd.Description,
		upd.Price,
		upd.Status,
		upd.Responsible,
	); errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23502":
			return model.Lead{}, model.ErrStatusNotExists
		default:
			return model.Lead{}, errors.Join(model.ErrDatabaseError, err)
		}
	} else if e.RowsAffected() == 0 {
		return model.Lead{}, model.ErrLeadNotExists
	} else {
		return l.GetLeadById(ctx, id)
	}
}
