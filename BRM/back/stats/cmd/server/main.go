package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"stats/cmd/server/factory"
	"stats/internal/app"
	"stats/internal/ports/grpcserver"
	"stats/internal/repo"
	adsRepo "stats/internal/repo/ads_repo"
	"stats/internal/repo/cache"
	coreRepo "stats/internal/repo/core_repo"
	leadsRepo "stats/internal/repo/leads_repo"
	"stats/pkg/logger"
	"syscall"
	"time"
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

	adsConn, err := factory.ConnectToPostgresAds(ctx)
	if err != nil {
		logs.Fatal(nil, err.Error())
	}
	defer func() {
		if adsConn != nil {
			adsConn.Close()
		}
	}()

	coreConn, err := factory.ConnectToPostgresCore(ctx)
	if err != nil {
		logs.Fatal(nil, err.Error())
	}
	defer func() {
		if coreConn != nil {
			coreConn.Close()
		}
	}()

	leadsConn, err := factory.ConnectToPostgresLeads(ctx)
	if err != nil {
		logs.Fatal(nil, err.Error())
	}
	defer func() {
		if leadsConn != nil {
			leadsConn.Close()
		}
	}()

	redisClient := factory.ConnectToRedis()
	if redisClient == nil {
		logs.Fatal(nil, "unable to connect to redis")
	}
	defer func() {
		if redisClient != nil {
			_ = redisClient.Close()
		}
	}()

	cacheRepo, err := cache.New(
		redisClient,
		time.Duration(viper.GetInt("redis-stats.expiration")),
	)
	if err != nil {
		logs.Fatal(nil, err.Error())
	}

	a := app.New(
		repo.New(
			adsRepo.New(adsConn),
			coreRepo.New(coreConn),
			leadsRepo.New(leadsConn),
			cacheRepo,
		), logs,
	)

	srv := grpcserver.New(a, logs)
	lis, err := factory.PrepareListener()
	if err != nil {
		logs.Fatal(nil, err.Error())
	}

	go func() {
		if err = srv.Serve(lis); err != nil {
			logs.Fatal(nil, fmt.Sprintf("starting grpc server: %s", err.Error()))
		}
	}()

	logs.Info(nil, "service stats successfully started")

	// preparing graceful shutdown
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT)

	// waiting for Ctrl+C
	<-osSignals
	srv.GracefulStop()
}
