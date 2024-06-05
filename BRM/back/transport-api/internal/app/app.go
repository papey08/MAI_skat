package app

import (
	"transport-api/internal/adapters/grpcads"
	"transport-api/internal/adapters/grpccore"
	"transport-api/internal/adapters/grpcleads"
	"transport-api/internal/adapters/grpcnotifications"
	"transport-api/internal/adapters/grpcstats"
)

type appImpl struct {
	core          grpccore.CoreClient
	ads           grpcads.AdsClient
	leads         grpcleads.LeadsClient
	stats         grpcstats.StatsClient
	notifications grpcnotifications.NotificationsClient
}

func NewApp(
	coreCli grpccore.CoreClient,
	adsCli grpcads.AdsClient,
	leadsCli grpcleads.LeadsClient,
	statsCli grpcstats.StatsClient,
	notificationsCli grpcnotifications.NotificationsClient,
) App {
	return &appImpl{
		core:          coreCli,
		ads:           adsCli,
		leads:         leadsCli,
		stats:         statsCli,
		notifications: notificationsCli,
	}
}
