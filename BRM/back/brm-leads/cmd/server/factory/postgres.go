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

func ConnectToPostgres(ctx context.Context) (*pgxpool.Pool, error) {
	adsRepoUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("postgres-leads.username"),
		os.Getenv("POSTGRES_LEADS_PASSWORD"),
		viper.GetString("postgres-leads.host"),
		viper.GetInt("postgres-leads.port"),
		viper.GetString("postgres-leads.dbname"),
		viper.GetString("postgres-leads.sslmode"))

	// 30 attempts to connect to postgres starting in docker container
	for i := 0; i < 30; i++ {
		conn, err := pgxpool.New(ctx, adsRepoUrl)
		if err != nil {
			time.Sleep(time.Second)
		} else {
			return conn, nil
		}
	}

	return nil, errors.New("unable to connect to postgres ads repo")
}
