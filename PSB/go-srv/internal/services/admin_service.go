package services

import (
	"context"
	"go-srv/internal/entities"
	"go-srv/internal/repo"
)

type AdminService struct {
	repo *repo.Repo
}

func NewAdminService(repo *repo.Repo) *AdminService {
	return &AdminService{
		repo: repo,
	}
}

func (a *AdminService) GetResponses(ctx context.Context, limit int, offset int) ([]entities.Response, error) {
	return a.repo.GetResponses(ctx, limit, offset)
}

func (a *AdminService) UpdateResponse(ctx context.Context, id int, category string) error {
	return a.repo.UpdateResponse(ctx, id, category)
}

func (a *AdminService) GetStatistics(ctx context.Context) (entities.Statistics, error) {
	return a.repo.GetStatistics(ctx)
}
