package ads_repo

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"stats/internal/model"
)

type adsRepoImpl struct {
	*pgxpool.Pool
}

const getActiveAdsAmountQuery = `
	SELECT COUNT(*) FROM "ads"
	WHERE "company_id" = $1 AND NOT "is_deleted";`

func (a *adsRepoImpl) GetActiveAdsAmount(ctx context.Context, companyId uint64) (uint, error) {
	row := a.QueryRow(ctx, getActiveAdsAmountQuery, companyId)
	var amount uint
	if err := row.Scan(&amount); err != nil {
		return 0, errors.Join(model.ErrAdsDatabase, err)
	}
	return amount, nil
}
