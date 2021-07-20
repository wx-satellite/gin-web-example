package middlewares

import (
	"errors"
	"gin-web/controller"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

func RateLimit(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就中断本次请求返回 rate limit...
		if bucket.TakeAvailable(1) < 1 {
			controller.SendFailResponse(c, errors.New("rate limit"))
			c.Abort()
			return
		}
		c.Next()
	}
}
