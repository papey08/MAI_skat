package repo

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type leadRepoImpl struct {
	*pgxpool.Pool
}
