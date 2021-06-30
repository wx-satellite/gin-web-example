package logger

import (
	"gin-web/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger 初始化Logger
func Init(cfg config.LogConfig) (err error) {
	writeSyncer := getLogWriter(
		cfg.Filename,
		cfg.MaxSize,
		cfg.MaxBackups,
		cfg.MaxAge,
	)
	encoder := getEncoder()

	// 将字符串转成 zap.InfoLevel 等内部值
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}

	core := zapcore.NewCore(encoder, writeSyncer, l)

	// zap.AddCaller 将调用函数信息记录到日志中的功能
	lg := zap.New(core, zap.AddCaller())

	zap.ReplaceGlobals(lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可

	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	// 序列化的参数设置
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,  // 文件名
		MaxSize:    maxSize,   // 切割之前日志的最大大小
		MaxBackups: maxBackup, // 保留的旧文件最大个数
		MaxAge:     maxAge,    // 保留旧文件的最大天数
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Close() {
	// 缓冲区日志flash
	_ = zap.L().Sync()
}
