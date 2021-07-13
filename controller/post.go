package controller

import (
	"gin-web/logic"
	"gin-web/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePost(ctx *gin.Context) {

	// 请求参数校验
	in := new(request.CreatePostRequest)
	if err := ctx.ShouldBindJSON(in); err != nil {
		SendFailResponse(ctx, err)
		return
	}

	// 填充 AuthorId

	// 业务逻辑处理
	if err := logic.CreatePost(in); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		SendFailResponse(ctx, err)
		return
	}
	// 返回
	SendSuccessResponse(ctx, nil)
}
