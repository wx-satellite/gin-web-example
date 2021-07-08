package middlewares

import (
	"errors"
	"fmt"
	"gin-web/controller"
	"gin-web/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authToken := c.Request.Header.Get("Authorization")
		_, _ = fmt.Sscanf(authToken, "Bearer %s", &authToken)
		if authToken == "" {
			controller.SendFailResponse(c, errors.New("请携带token信息！"))
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(authToken)
		if err != nil {
			controller.SendFailResponse(c, errors.New("token非法！"))
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("userId", mc.UserId)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
