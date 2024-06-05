package app

import (
	"context"
	"registration/internal/adapters/grpccore"
	"registration/internal/model"
)

type appImpl struct {
	core grpccore.CoreClient
}

func NewApp(coreCli grpccore.CoreClient) App {
	return &appImpl{core: coreCli}
}

func (a *appImpl) CreateCompanyAndOwner(ctx context.Context, company model.Company, owner model.Employee) (model.Company, model.Employee, error) {
	return a.core.CreateCompanyAndOwner(ctx, company, owner)
}

func (a *appImpl) GetIndustriesList(ctx context.Context) (map[string]uint64, error) {
	return a.core.GetIndustriesList(ctx)
}
