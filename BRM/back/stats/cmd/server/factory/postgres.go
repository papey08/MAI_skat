package factory

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	"os"
	"time"
)

func ConnectToPostgresAds(ctx context.Context) (*pgxpool.Pool, error) {
	coreRepoUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("postgres-ads.username"),
		os.Getenv("POSTGRES_ADS_PASSWORD"),
		viper.GetString("postgres-ads.host"),
		viper.GetInt("postgres-ads.port"),
		viper.GetString("postgres-ads.dbname"),
		viper.GetString("postgres-ads.sslmode"))

	// 30 attempts to connect to postgres starting in docker container
	for i := 0; i < 30; i++ {
		conn, err := pgxpool.New(ctx, coreRepoUrl)
		if err != nil {
			time.Sleep(time.Second)
		} else {
			return conn, nil
		}
	}

	return nil, errors.New("unable to connect to postgres ads repo")
}

func ConnectToPostgresCore(ctx context.Context) (*pgxpool.Pool, error) {
	coreRepoUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("postgres-core.username"),
		os.Getenv("POSTGRES_CORE_PASSWORD"),
		viper.GetString("postgres-core.host"),
		viper.GetInt("postgres-core.port"),
		viper.GetString("postgres-core.dbname"),
		viper.GetString("postgres-core.sslmode"))

	// 30 attempts to connect to postgres starting in docker container
	for i := 0; i < 30; i++ {
		conn, err := pgxpool.New(ctx, coreRepoUrl)
		if err != nil {
			time.Sleep(time.Second)
		} else {
			return conn, nil
		}
	}

	return nil, errors.New("unable to connect to postgres core repo")
}

func ConnectToPostgresLeads(ctx context.Context) (*pgxpool.Pool, error) {
	coreRepoUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("postgres-leads.username"),
		os.Getenv("POSTGRES_LEADS_PASSWORD"),
		viper.GetString("postgres-leads.host"),
		viper.GetInt("postgres-leads.port"),
		viper.GetString("postgres-leads.dbname"),
		viper.GetString("postgres-leads.sslmode"))

	// 30 attempts to connect to postgres starting in docker container
	for i := 0; i < 30; i++ {
		conn, err := pgxpool.New(ctx, coreRepoUrl)
		if err != nil {
			time.Sleep(time.Second)
		} else {
			return conn, nil
		}
	}

	return nil, errors.New("unable to connect to postgres leads repo")
}
