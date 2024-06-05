package cache

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"stats/internal/model"
	"time"
)

type Cache interface {
	AddCompanyData(ctx context.Context, companyId uint64, data model.MainPageStats)
	GetCompanyData(ctx context.Context, companyId uint64) (model.MainPageStats, bool)
}

func New(cli *redis.Client, expirationTime time.Duration) (Cache, error) {
	if cli == nil {
		return nil, errors.New("unable to create cache")
	}
	return &cacheImpl{
		expirationTime: expirationTime,
		Client:         cli,
	}, nil
}
