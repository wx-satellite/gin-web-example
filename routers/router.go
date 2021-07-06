package routers

import (
	"gin-web/controller"
	"gin-web/routers/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(engine *gin.Engine, mw ...gin.HandlerFunc) {
	engine.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))

	engine.Use(mw...)

	// 注册
	engine.POST("/sign_up", controller.SignUpHandler)
	// 登陆
	engine.POST("/sign_in", controller.SignInHandler)

	// 心跳检测
	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 为没有配置处理函数的路由添加处理函数
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "404",
		})
	})

}
