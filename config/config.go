package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Instance = new(AppConfig)

// 配置文件设置为结构体的目的是为了更直观的看到各个模块的配置信息
type AppConfig struct {
	Name        string `mapstructure:"name"`
	Mode        string `mapstructure:"mode"`
	Version     string `mapstructure:"version"`
	Port        int    `mapstructure:"port"`
	StartTime   string `mapstructure:"start_time"`
	MachineID   int64  `mapstructure:"machine_id"`
	LogConfig   `mapstructure:"log"`
	MysqlConfig `mapstructure:"mysql"`
	RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
	MaxOpen  int    `mapstructure:"max_open"`
	MaxIdle  int    `mapstructure:"max_idle "`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init(filename string) (err error) {
	if filename != "" {
		viper.SetConfigFile(filename)
	} else {
		viper.SetConfigName("app") // 文件名称（ 不需要指定后缀名 ）
		viper.AddConfigPath(".")   // 去指定的路径里搜索目标文件，可以多次设值
	}
	if err = viper.ReadInConfig(); nil != err {
		return
	}
	// 将配置文件的信息反序列化到 Conf 变量中
	if err = viper.Unmarshal(Instance); nil != err {
		return
	}
	// 热加载
	viper.WatchConfig()
	// 配置发现变更时的回调函数
	viper.OnConfigChange(func(in fsnotify.Event) {
		c := new(AppConfig)
		if err = viper.Unmarshal(c); nil != err {
			return
		}
		Instance = c
	})
	return
}
