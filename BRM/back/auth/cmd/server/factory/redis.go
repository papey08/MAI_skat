package factory

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"os"
)

func ConnectToRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis-tokens.host"),
			viper.GetInt("redis-tokens.port")),
		Password: os.Getenv("REDIS_TOKENS_PASSWORD"),
		DB:       viper.GetInt("redis-tokens.db"),
	})
}
