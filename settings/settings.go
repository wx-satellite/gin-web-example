package settings

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	if err = viper.ReadInConfig(); nil != err {
		return
	}
	// 热加载
	viper.WatchConfig()
	// 配置发现变更时的回调函数
	viper.OnConfigChange(func(in fsnotify.Event) {

	})
	return
}
