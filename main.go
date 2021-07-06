package main

import (
	"context"
	"flag"
	"fmt"
	"gin-web/routers"
	"gin-web/sysinit"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 指定配置文件
	var filename string
	flag.StringVar(&filename, "f", "", "配置文件")
	flag.Parse()
	sysinit.Init(filename)
	gin.SetMode(gin.ReleaseMode)

	g := gin.New()

	// 加载路由
	routers.Load(g)

	srv := &http.Server{
		//Addr:    fmt.Sprintf(":%d", config.Instance.Port),
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: g,
	}

	// 开启web服务
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sysinit.Close()
			// 终止程序
			zap.L().Fatal("服务启动失败", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("关闭服务失败", zap.Error(err))
	}

	// 关闭资源
	sysinit.Close()
}
