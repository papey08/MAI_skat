package main

import (
	"brm-ads/cmd/server/factory"
	"brm-ads/internal/adapters/grpccore"
	"brm-ads/internal/adapters/grpcleads"
	"brm-ads/internal/app"
	"brm-ads/internal/ports/grpcserver"
	"brm-ads/internal/repo"
	"brm-ads/pkg/logger"
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

	adsRepo, err := factory.ConnectToPostgres(ctx)
	if err != nil {
		logs.Fatal(nil, err.Error())
	}
	defer func() {
		if adsRepo != nil {
			adsRepo.Close()
		}
	}()

	coreClient, err := grpccore.NewCoreClient(ctx, fmt.Sprintf("%s:%d",
		viper.GetString("grpc-clients.core.host"),
		viper.GetInt("grpc-clients.core.port")))
	if err != nil {
		logs.Fatal(nil, fmt.Sprintf("create grpc core client: %s", err.Error()))
	}

	leadsClient, err := grpcleads.NewLeadsClient(ctx, fmt.Sprintf("%s:%d",
		viper.GetString("grpc-clients.leads.host"),
		viper.GetInt("grpc-clients.leads.port")))
	if err != nil {
		logs.Fatal(nil, fmt.Sprintf("create grpc leads client: %s", err.Error()))
	}

	a := app.New(repo.New(adsRepo), coreClient, leadsClient, logs)

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

	logs.Info(nil, "service brm-ads successfully started")
	<-osSignals
	srv.GracefulStop()
}
