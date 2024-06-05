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
			viper.GetString("redis-stats.host"),
			viper.GetInt("redis-stats.port")),
		Password: os.Getenv("REDIS_STATS_PASSWORD"),
		DB:       viper.GetInt("redis-stats.db"),
	})
}
