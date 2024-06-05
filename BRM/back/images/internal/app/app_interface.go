package app

import (
	"context"
	"images/internal/model"
	"images/internal/repo"
	"images/pkg/logger"
)

type App interface {
	AddImage(ctx context.Context, img model.Image) (uint64, error)
	GetImage(ctx context.Context, id uint64) (model.Image, error)
}

func New(r repo.ImageRepo, logs logger.Logger) App {
	return &appImpl{
		r:    r,
		logs: logs,
	}
}
