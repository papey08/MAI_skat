package repo

import (
	"brm-ads/internal/model"
	"context"
	"errors"
)

const (
	createResponseQuery = `
		INSERT INTO "responses" ("company_id", "employee_id", "ad_id", "creation_date")
		VALUES ($1, $2, $3, $4)
		RETURNING "id";`

	getResponsesQuery = `
		SELECT * FROM "responses"
		WHERE "company_id" = $1
		LIMIT $2 OFFSET $3;`

	getResponsesAmountQuery = `
		SELECT COUNT(*) FROM "responses"
		WHERE "company_id" = $1;`
)

func (a *adRepoImpl) CreateResponse(ctx context.Context, resp model.Response) (model.Response, error) {
	var respId uint64
	if err := a.QueryRow(ctx, createResponseQuery,
		resp.CompanyId,
		resp.EmployeeId,
		resp.AdId,
		resp.CreationDate,
	).Scan(&respId); err != nil {
		return model.Response{}, errors.Join(model.ErrDatabaseError, err)
	} else {
		resp.Id = respId
		return resp, nil
	}
}

func (a *adRepoImpl) GetResponses(ctx context.Context, companyId uint64, limit uint, offset uint) ([]model.Response, uint, error) {
	var amount uint
	if err := a.QueryRow(ctx, getResponsesAmountQuery,
		companyId,
	).Scan(&amount); err != nil {
		return []model.Response{}, 0, errors.Join(model.ErrDatabaseError, err)
	} else if amount == 0 {
		return []model.Response{}, 0, nil
	}

	rows, err := a.Query(ctx, getResponsesQuery,
		companyId,
		limit,
		offset)
	if err != nil {
		return []model.Response{}, 0, errors.Join(model.ErrDatabaseError, err)
	}
	defer rows.Close()

	responses := make([]model.Response, 0)
	for rows.Next() {
		var resp model.Response
		_ = rows.Scan(
			&resp.Id,
			&resp.CompanyId,
			&resp.EmployeeId,
			&resp.AdId,
			&resp.CreationDate,
		)
		responses = append(responses, resp)
	}
	return responses, amount, nil
}
