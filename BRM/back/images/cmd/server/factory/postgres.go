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
	imagesRepoUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("postgres-images.username"),
		os.Getenv("POSTGRES_IMAGES_PASSWORD"),
		viper.GetString("postgres-images.host"),
		viper.GetInt("postgres-images.port"),
		viper.GetString("postgres-images.dbname"),
		viper.GetString("postgres-images.sslmode"))

	// 30 attempts to connect to postgres starting in docker container
	for i := 0; i < 30; i++ {
		conn, err := pgxpool.New(ctx, imagesRepoUrl)
		if err != nil {
			time.Sleep(time.Second)
		} else {
			return conn, nil
		}
	}

	return nil, errors.New("unable to connect to postgres images repo")
}
