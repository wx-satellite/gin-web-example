package sysinit

import (
	"gin-web/config"
	"gin-web/dao/mysql"
	"gin-web/dao/redis"
	"gin-web/logger"
	"gin-web/pkg/snowflake"
)

func Init(filename string) {
	if err := config.Init(filename); nil != err {
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
	if err := snowflake.Init(config.Instance.StartTime, config.Instance.MachineID); nil != err {
		panic(err)
	}
}

func Close() {
	logger.Close()
	mysql.Close()
	redis.Close()
}
