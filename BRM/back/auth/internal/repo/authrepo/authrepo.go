package authrepo

import (
	"auth/internal/model"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type authRepoImpl struct {
	refreshExpiration time.Duration
	*redis.Client
}

func (a *authRepoImpl) SetTokens(ctx context.Context, tokens model.TokensPair) error {
	if err := a.Set(ctx, tokens.Access, tokens.Refresh, a.refreshExpiration).Err(); err != nil {
		return model.ErrAuthRepoError
	}
	return nil
}

func (a *authRepoImpl) GetTokens(ctx context.Context, access string) (model.TokensPair, error) {
	if refresh, err := a.Get(ctx, access).Result(); errors.Is(err, redis.Nil) {
		return model.TokensPair{}, model.ErrTokensNotExist
	} else if err != nil {
		return model.TokensPair{}, model.ErrAuthRepoError
	} else {
		return model.TokensPair{
			Access:  access,
			Refresh: refresh,
		}, nil
	}
}

func (a *authRepoImpl) DeleteTokens(ctx context.Context, access string) error {
	if _, err := a.Del(ctx, access).Result(); err == nil || errors.Is(err, redis.Nil) {
		return nil
	} else {
		return model.ErrAuthRepoError
	}
}
