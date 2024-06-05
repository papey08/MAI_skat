package app

import (
	"context"
	"notifications/internal/adapters/grpcstats"
	"notifications/internal/model"
	"notifications/internal/repo"
	"notifications/pkg/logger"
)

type App interface {
	CreateNotification(ctx context.Context, notification model.Notification) error

	GetNotifications(ctx context.Context, companyId uint64, limit uint, offset uint, onlyNotViewed bool) ([]model.Notification, uint, error)
	GetNotification(ctx context.Context, companyId uint64, notificationId uint64) (model.Notification, error)
	SubmitClosedLead(ctx context.Context, companyId uint64, notificationId uint64, submit bool) error
}

func New(r repo.NotificationsRepo, cli grpcstats.StatsClient, logs logger.Logger) App {
	return &appImpl{
		r:        r,
		statsCli: cli,
		logs:     logs,
	}
}
