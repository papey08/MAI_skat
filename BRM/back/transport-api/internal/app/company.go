package app

import (
	"context"
	"transport-api/internal/model/core"
)

func (a *appImpl) GetCompany(ctx context.Context, id uint64) (core.Company, error) {
	return a.core.GetCompany(ctx, id)
}

func (a *appImpl) UpdateCompany(ctx context.Context, companyId uint64, ownerId uint64, upd core.UpdateCompany) (core.Company, error) {
	return a.core.UpdateCompany(ctx, companyId, ownerId, upd)
}

func (a *appImpl) GetCompanyIndustries(ctx context.Context) (map[string]uint64, error) {
	return a.core.GetIndustries(ctx)
}
