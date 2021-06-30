package sysinit

import (
	"gin-web/config"
	"gin-web/dao/mysql"
	"gin-web/dao/redis"
	"gin-web/logger"
)

func init() {
	if err := config.Init(); nil != err {
		panic(err)
	}
	if err := logger.Init(config.Instance.LogConfig); nil != err {
		panic(err)
	}
	if err := mysql.Init(config.Instance.MysqlConfig); nil != err {
		panic(err)
		return
	}
	if err := redis.Init(config.Instance.RedisConfig); nil != err {
		panic(err)
	}
}

func Close() {
	logger.Close()
	mysql.Close()
	redis.Close()
}
