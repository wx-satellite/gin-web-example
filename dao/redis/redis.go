package redis

import (
	"context"
	"fmt"
	"gin-web/config"
	"github.com/go-redis/redis/v8"
	"time"
)

var rdb *redis.Client

func Init(cfg config.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.Db,       // use default DB
		PoolSize: cfg.PoolSize, // 连接池大小
	})
	// 超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = rdb.Ping(ctx).Result()
	return
}

func Close() {
	_ = rdb.Close()
}
