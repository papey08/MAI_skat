package passrepo

import (
	"auth/internal/model"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PassRepo interface {
	CreateEmployee(ctx context.Context, employee model.Employee) error
	GetEmployee(ctx context.Context, email string) (model.Employee, error)
	DeleteEmployee(ctx context.Context, email string) error
}

func New(pool *pgxpool.Pool) PassRepo {
	return &passRepoImpl{
		Pool: pool,
	}
}
