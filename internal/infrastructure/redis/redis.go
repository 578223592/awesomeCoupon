package redis

import (
	"github.com/578223592/awesomeCoupon/internal/infrastructure/handler/config"
	"sync"

	"github.com/redis/go-redis/v9"
)

var initOnce sync.Once
var client *redis.Client

func Init() {
	initOnce.Do(func() {
		rdb := redis.NewClient(&redis.Options{
			Addr:     config.Viper.GetString("redis.address"),
			Password: config.Viper.GetString("redis.password"),
			DB:       config.Viper.GetInt("redis.db"),
		})
		client = rdb
	})

}

func GetClient() *redis.Client {
	Init()
	return client
}
