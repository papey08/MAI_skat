package main

import (
	"brm-core/cmd/server/factory"
	"brm-core/internal/adapters/grpcauth"
	"brm-core/internal/app"
	"brm-core/internal/ports/grpcserver"
	"brm-core/internal/repo"
	"brm-core/pkg/logger"
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

	coreRepo, err := factory.ConnectToPostgres(ctx)
	if err != nil {
		logs.Fatal(nil, err.Error())
	}
	defer func() {
		if coreRepo != nil {
			coreRepo.Close()
		}
	}()

	authClient, err := grpcauth.NewAuthClient(ctx, fmt.Sprintf("%s:%d",
		viper.GetString("grpc-auth-client.host"),
		viper.GetInt("grpc-auth-client.port")))

	a := app.New(repo.New(coreRepo), authClient, logs)

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

	logs.Info(nil, "service brm-core successfully started")
	<-osSignals
	srv.GracefulStop()
}
