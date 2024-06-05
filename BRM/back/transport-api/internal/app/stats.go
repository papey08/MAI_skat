package app

import (
	"context"
	"transport-api/internal/model/stats"
)

func (a *appImpl) GetCompanyMainPage(ctx context.Context, companyId uint64) (stats.MainPage, error) {
	return a.stats.GetCompanyMainPage(ctx, companyId)
}
