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
	coreRepoUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		viper.GetString("postgres-passwords.username"),
		os.Getenv("POSTGRES_PASSWORDS_PASSWORD"),
		viper.GetString("postgres-passwords.host"),
		viper.GetInt("postgres-passwords.port"),
		viper.GetString("postgres-passwords.dbname"),
		viper.GetString("postgres-passwords.sslmode"))

	// 30 attempts to connect to postgres starting in docker container
	for i := 0; i < 30; i++ {
		conn, err := pgxpool.New(ctx, coreRepoUrl)
		if err != nil {
			time.Sleep(time.Second)
		} else {
			return conn, nil
		}
	}

	return nil, errors.New("unable to connect to postgres passwords repo")
}
