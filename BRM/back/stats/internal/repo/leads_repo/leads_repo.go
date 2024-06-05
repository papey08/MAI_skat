package leads_repo

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"stats/internal/model"
)

type leadsRepoImpl struct {
	*pgxpool.Pool
}

const (
	getActiveLeadsStatsQuery = `
		SELECT COUNT(*), SUM("price") FROM "leads"
		WHERE "company_id" = $1 AND NOT "is_deleted" AND "status" BETWEEN 1 AND 4;`

	getClosedLeadsStatsQuery = `
		SELECT COUNT(*), SUM("price") FROM "leads"
		WHERE "company_id" = $1 AND NOT "is_deleted" AND "status" = 5;`
)

func (l *leadsRepoImpl) GetMainPageLeadsStats(ctx context.Context, companyId uint64) (model.MainPageStats, error) {
	var data model.MainPageStats
	var pgErr pgx.ScanArgError
	row := l.QueryRow(ctx, getActiveLeadsStatsQuery, companyId)
	if err := row.Scan(&data.ActiveLeadsAmount, &data.ActiveLeadsPrice); errors.As(err, &pgErr) {
		data.ActiveLeadsPrice = 0.
	} else if err != nil {
		return model.MainPageStats{}, errors.Join(model.ErrLeadsDatabase, err)
	}

	row = l.QueryRow(ctx, getClosedLeadsStatsQuery, companyId)
	if err := row.Scan(&data.ClosedLeadsAmount, &data.ClosedLeadsPrice); errors.As(err, &pgErr) {
		data.ClosedLeadsPrice = 0.
	} else if err != nil {
		return model.MainPageStats{}, errors.Join(model.ErrLeadsDatabase, err)
	}

	return data, nil
}
