package main

import (
	"auth/cmd/server/factory"
	"auth/internal/app"
	"auth/internal/app/tokenizer"
	"auth/internal/ports/grpcserver"
	"auth/internal/ports/httpserver"
	"auth/internal/repo/authrepo"
	"auth/internal/repo/passrepo"
	"auth/pkg/logger"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//	@title			BRM API
//	@version		1.0
//	@description	Swagger документация к API авторизации
//	@host			localhost:8092
//	@BasePath		/api/v1/auth

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

	passConn, err := factory.ConnectToPostgres(ctx)
	if err != nil {
		logs.Fatal(nil, err.Error())
	}
	defer func() {
		if passConn != nil {
			passConn.Close()
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

	authRepo, err := authrepo.New(
		redisClient,
		time.Duration(viper.GetInt("app.refresh-token-expiration")))
	if err != nil {
		logs.Fatal(nil, err.Error())
	}

	passRepo := passrepo.New(passConn)

	tkn := tokenizer.New(
		time.Duration(viper.GetInt("app.access-token-expiration")),
		[]byte(os.Getenv("SIGNKEY")),
	)

	a := app.New(
		os.Getenv("PASSWORDSALT"),
		viper.GetInt("app.refresh-token-length"),
		authRepo,
		passRepo,
		tkn,
		logs,
	)

	grpcsrv := grpcserver.New(a, logs)
	lis, err := factory.PrepareListener()
	if err != nil {
		logs.Fatal(nil, err.Error())
	}

	go func() {
		if err = grpcsrv.Serve(lis); err != nil {
			logs.Fatal(nil, fmt.Sprintf("starting grpc server: %s", err.Error()))
		}
	}()

	httpsrv := httpserver.New(
		fmt.Sprintf("%s:%d",
			viper.GetString("http-server.host"),
			viper.GetInt("http-server.port")),
		viper.GetStringSlice("origins"),
		a, logs)

	go func() {
		if err = httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logs.Fatal(nil, fmt.Sprintf("listening server: %s", err.Error()))
		}
	}()

	logs.Info(nil, "service auth successfully started")

	// preparing graceful shutdown
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT)

	// waiting for Ctrl+C
	<-osSignals

	shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second) // 30s timeout to finish all active connections
	defer cancel()

	grpcsrv.GracefulStop()
	_ = httpsrv.Shutdown(shutdownCtx)
}
