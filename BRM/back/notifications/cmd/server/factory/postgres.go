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
	notificationsRepoUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("postgres-notifications.username"),
		os.Getenv("POSTGRES_NOTIFICATIONS_PASSWORD"),
		viper.GetString("postgres-notifications.host"),
		viper.GetInt("postgres-notifications.port"),
		viper.GetString("postgres-notifications.dbname"),
		viper.GetString("postgres-notifications.sslmode"))

	// 30 attempts to connect to postgres starting in docker container
	for i := 0; i < 30; i++ {
		pool, err := pgxpool.New(ctx, notificationsRepoUrl)
		if err != nil {
			time.Sleep(time.Second)
		} else {
			return pool, nil
		}
	}

	return nil, errors.New("unable to connect to postgres notifications repo")
}
