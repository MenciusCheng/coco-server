package db

import (
	"coco-server/conf"
	ulog "coco-server/util/log"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var RedisCon *redis.Client

func InitRedis(ctx context.Context) {
	redisConfigs := conf.Conf.RedisConfigs
	if len(redisConfigs) == 0 {
		err := errors.New("redisConfigs is empty")
		ulog.Error(ctx, "InitRedis err", zap.Error(err))
		panic(err)
	}

	redisConfig := redisConfigs[0]
	con := redis.NewClient(&redis.Options{
		PoolSize: 100,
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.Database,
	})
	_, err := con.Ping(ctx).Result()
	if err != nil {
		panic("redis初始化失败, err:" + err.Error())
	}

	RedisCon = con
}
