package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"images/cmd/server/factory"
	"images/internal/app"
	"images/internal/ports/httpserver"
	"images/internal/repo"
	"images/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	dockerConfigFile = "config/config-docker.yml"
	localConfigFile  = "config/config-local.yml"
)

//	@title			BRM API
//	@version		1.0
//	@description	Swagger документация к API сервиса для хранения изображений
//	@host			localhost:8093
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
	if err := godotenv.Load("config/.env"); err != nil {
		logs.Fatal(nil, fmt.Sprintf("unable to load .env file: %s", err.Error()))
	}

	imagesRepo, err := factory.ConnectToPostgres(ctx)
	if err != nil {
		logs.Fatal(nil, fmt.Sprintf("connect to images databse: %s", err.Error()))
	}
	defer imagesRepo.Close()

	a := app.New(repo.New(imagesRepo), logs)

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

	logs.Info(nil, "service images successfully started")

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
