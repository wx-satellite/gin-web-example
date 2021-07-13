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
	v1 := engine.Group("/api/v1")
	// 注册
	v1.POST("/sign_up", controller.SignUpHandler)
	// 登陆
	v1.POST("/sign_in", controller.SignInHandler)

	// 心跳检测
	v1.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 社区
	v1.Use(middlewares.JWTAuth())
	{
		v1.GET("/community", controller.Communities)
		v1.GET("/community/:id", controller.CommunityDetail)

		// 帖子
		v1.POST("/post", controller.CreatePost)
	}

	// 为没有配置处理函数的路由添加处理函数
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "404",
		})
	})

}
