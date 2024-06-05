package app

import (
	"auth/internal/app/tokenizer"
	"auth/internal/model"
	"auth/internal/repo/authrepo"
	"auth/internal/repo/passrepo"
	"auth/pkg/logger"
	"context"
)

type App interface {
	RegisterEmployee(ctx context.Context, employee model.Employee) error
	DeleteEmployee(ctx context.Context, email string) error

	LoginEmployee(ctx context.Context, email string, password string) (model.TokensPair, error)
	RefreshTokens(ctx context.Context, tokens model.TokensPair) (model.TokensPair, error)
	LogoutEmployee(ctx context.Context, tokens model.TokensPair) error
}

func New(
	passwordSalt string,
	refreshTokenLength int,
	authRepo authrepo.AuthRepo,
	passRepo passrepo.PassRepo,
	tokenizer tokenizer.Tokenizer,
	logs logger.Logger,
) App {
	return &appImpl{
		passwordSalt:       passwordSalt,
		refreshTokenLength: refreshTokenLength,
		authRepo:           authRepo,
		passRepo:           passRepo,
		tokenizer:          tokenizer,
		logs:               logs,
	}
}
