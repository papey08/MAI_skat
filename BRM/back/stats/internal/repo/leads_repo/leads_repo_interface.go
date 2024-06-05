package leads_repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"stats/internal/model"
)

type LeadsRepo interface {
	GetMainPageLeadsStats(ctx context.Context, companyId uint64) (model.MainPageStats, error)
}

func New(pool *pgxpool.Pool) LeadsRepo {
	return &leadsRepoImpl{
		Pool: pool,
	}
}
