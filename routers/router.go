package routers

import (
	"gin-web/routers/middlewares"
	"github.com/gin-gonic/gin"
)

func Load(engine *gin.Engine, mw ...gin.HandlerFunc) {
	engine.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))

	engine.Use(mw...)

}
