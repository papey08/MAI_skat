package repo

import (
	"brm-leads/internal/model"
	"context"
)

const (
	getStatusesQuery = `
		SELECT * FROM "statuses";`
)

func (l *leadRepoImpl) GetStatuses(ctx context.Context) (map[string]uint64, error) {
	rows, err := l.Query(ctx, getStatusesQuery)
	if err != nil {
		return map[string]uint64{}, model.ErrDatabaseError
	}
	defer rows.Close()

	statuses := make(map[string]uint64)
	for rows.Next() {
		var id uint64
		var status string
		_ = rows.Scan(&id, &status)
		statuses[status] = id
	}
	return statuses, nil
}
