package repo

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type coreRepoImpl struct {
	*pgxpool.Pool
}
