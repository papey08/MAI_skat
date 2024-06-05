package repo

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type adRepoImpl struct {
	*pgxpool.Pool
}
