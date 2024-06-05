package repo

import (
	"brm-leads/internal/model"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LeadsRepo interface {
	GetLeadById(ctx context.Context, id uint64) (model.Lead, error)
	GetLeads(ctx context.Context, companyId uint64, filter model.Filter) ([]model.Lead, uint, error)
	CreateLead(ctx context.Context, lead model.Lead) (model.Lead, error)
	UpdateLead(ctx context.Context, id uint64, upd model.UpdateLead) (model.Lead, error)

	GetStatuses(ctx context.Context) (map[string]uint64, error)
}

func New(pool *pgxpool.Pool) LeadsRepo {
	return &leadRepoImpl{
		Pool: pool,
	}
}
