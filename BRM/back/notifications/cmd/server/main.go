package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"notifications/cmd/server/factory"
	"notifications/internal/adapters/grpcstats"
	"notifications/internal/app"
	"notifications/internal/ports/grpcserver"
	"notifications/internal/repo"
	"notifications/pkg/logger"
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

	notificationsRepo, err := factory.ConnectToPostgres(ctx)
	if err != nil {
		logs.Fatal(nil, err.Error())
	}
	defer notificationsRepo.Close()

	statsClient, err := grpcstats.NewStatsClient(ctx, fmt.Sprintf("%s:%d",
		viper.GetString("grpc-clients.stats.host"),
		viper.GetInt("grpc-clients.stats.port")))
	if err != nil {
		logs.Fatal(nil, fmt.Sprintf("create grpc stats client: %s", err.Error()))
	}

	a := app.New(repo.New(notificationsRepo), statsClient, logs)

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

	logs.Info(nil, "service notifications successfully started")
	<-osSignals
	srv.GracefulStop()
}
