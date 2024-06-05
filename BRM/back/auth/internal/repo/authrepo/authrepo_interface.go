package authrepo

import (
	"auth/internal/model"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type AuthRepo interface {
	SetTokens(ctx context.Context, tokens model.TokensPair) error
	GetTokens(ctx context.Context, access string) (model.TokensPair, error)
	DeleteTokens(ctx context.Context, access string) error
}

func New(cli *redis.Client, refreshExpiration time.Duration) (AuthRepo, error) {
	if cli == nil {
		return nil, errors.New("unable to create AuthRepo")
	}
	return &authRepoImpl{
		refreshExpiration: refreshExpiration,
		Client:            cli,
	}, nil
}
