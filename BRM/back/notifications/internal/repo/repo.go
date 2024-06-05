package repo

import "github.com/jackc/pgx/v5/pgxpool"

type notificationsRepoImpl struct {
	*pgxpool.Pool
}
