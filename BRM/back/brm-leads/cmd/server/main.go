package main

import (
	"brm-leads/cmd/server/factory"
	"brm-leads/internal/adapters/grpcads"
	"brm-leads/internal/adapters/grpccore"
	"brm-leads/internal/adapters/grpcnotifications"
	"brm-leads/internal/app"
	"brm-leads/internal/ports/grpcserver"
	"brm-leads/internal/repo"
	"brm-leads/pkg/logger"
	"context"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

const (
	dockerConfigFile = "config/config-docker.yml"
	localConfigFile  = "config/config-local.yml"
)

func main() {
	ctx := context.Background()
	logs := logger.New()

	isDocker := flag.Bool("docker", false, "flag if this project is running in docker container")
	flag.Parse()
	var configPath string
	if *isDocker {
		configPath = dockerConfigFile
	} else {
		configPath = localConfigFile
	}

	if err := factory.SetConfigs(configPath); err != nil {
		logs.Fatal(nil, err.Error())
	}

	leadsRepo, err := factory.ConnectToPostgres(ctx)
	if err != nil {
		logs.Fatal(nil, fmt.Sprintf("connect to leads databse: %s", err.Error()))
	}
	defer func() {
		leadsRepo.Close()
	}()

	coreClient, err := grpccore.NewCoreClient(ctx, fmt.Sprintf("%s:%d",
		viper.GetString("grpc-clients.core.host"),
		viper.GetInt("grpc-clients.core.port")))
	if err != nil {
		logs.Fatal(nil, fmt.Sprintf("create grpc core client: %s", err.Error()))
	}

	adsClient, err := grpcads.NewAdsClient(ctx, fmt.Sprintf("%s:%d",
		viper.GetString("grpc-clients.ads.host"),
		viper.GetInt("grpc-clients.ads.port")))
	if err != nil {
		logs.Fatal(nil, fmt.Sprintf("create grpc ads client: %s", err.Error()))
	}

	notificationsClient, err := grpcnotifications.NewNotificationsClient(ctx, fmt.Sprintf("%s:%d",
		viper.GetString("grpc-clients.notifications.host"),
		viper.GetInt("grpc-clients.notifications.port")))
	if err != nil {
		logs.Fatal(nil, fmt.Sprintf("create grpc notifications client: %s", err.Error()))
	}

	a := app.New(
		repo.New(leadsRepo),
		coreClient,
		adsClient,
		notificationsClient,
		logs,
	)

	srv := grpcserver.New(a, logs)
	lis, err := factory.PrepareListener()
	if err != nil {
		logs.Fatal(nil, err.Error())
	}

	// preparing graceful shutdown
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT)

	go func() {
		if err = srv.Serve(lis); err != nil {
			logs.Fatal(nil, fmt.Sprintf("starting grpc server: %s", err.Error()))
		}
	}()

	logs.Info(nil, "service brm-leads successfully started")
	<-osSignals
	srv.GracefulStop()
}
