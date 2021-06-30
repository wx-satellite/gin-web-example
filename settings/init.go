package settings

import (
	"gin-web/dao/mysql"
	"gin-web/dao/redis"
	"gin-web/logger"
)

func init() {
	if err := Init(); nil != err {
		panic(err)
	}
	if err := logger.Init(); nil != err {
		panic(err)
	}
	if err := mysql.Init(); nil != err {
		panic(err)
		return
	}
	if err := redis.Init(); nil != err {
		panic(err)
	}
}

func Close() {
	logger.Close()
	mysql.Close()
	redis.Close()
}
