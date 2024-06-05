package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"registration/internal/adapters/grpccore"
	"registration/internal/app"
	"registration/internal/ports/httpserver"
	"registration/pkg/logger"
	"syscall"
	"time"
)

const (
	dockerConfigFile = "config/config-docker.yml"
	localConfigFile  = "config/config-local.yml"
)

//	@title			BRM API
//	@version		1.0
//	@description	Swagger документация к API
//	@host			localhost:8091
//	@BasePath		/api/v1

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

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		logs.Fatal(nil, fmt.Sprintf("reading config: %s", err.Error()))
	}

	coreClient, err := grpccore.NewCoreClient(ctx, fmt.Sprintf("%s:%d",
		viper.GetString("grpc-clients.core.host"),
		viper.GetInt("grpc-clients.core.port")))
	if err != nil {
		logs.Fatal(nil, fmt.Sprintf("create grpc core client: %s", err.Error()))
	}

	a := app.NewApp(coreClient)

	srv := httpserver.New(
		fmt.Sprintf("%s:%d",
			viper.GetString("http-server.host"),
			viper.GetInt("http-server.port")),
		viper.GetStringSlice("origins"),
		a, logs)

	go func() {
		if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logs.Fatal(nil, fmt.Sprintf("listening server: %s", err.Error()))
		}
	}()

	logs.Info(nil, "service registration successfully started")

	// preparing graceful shutdown
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT)

	// waiting for Ctrl+C
	<-osSignals

	shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second) // 30s timeout to finish all active connections
	defer cancel()

	_ = srv.Shutdown(shutdownCtx)
}
