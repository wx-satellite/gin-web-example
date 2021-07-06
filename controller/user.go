package controller

import (
	"gin-web/logic"
	"gin-web/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(ctx *gin.Context) {
	// 参数校验
	in := new(request.SignUpRequest)
	if err := ctx.ShouldBindJSON(in); err != nil {
		zap.L().Error("SignUp with invalid params", zap.Error(err))
		SendFailResponse(ctx, err)
		return
	}
	// 业务逻辑处理
	if err := logic.SignUp(in); err != nil {
		zap.L().Error("SignUp fail", zap.Error(err))
		SendFailResponse(ctx, err)
		return
	}
	// 返回结果
	SendSuccessResponse(ctx, nil)
}

func SignInHandler(ctx *gin.Context) {
	// 参数校验
	in := new(request.SignInRequest)
	if err := ctx.ShouldBindJSON(in); err != nil {
		zap.L().Error("SignIn with invalid params", zap.Error(err))
		SendFailResponse(ctx, err)
		return
	}
	// 业务逻辑处理
	if err := logic.SignIn(in); err != nil {
		zap.L().Error("SignIn fail", zap.Error(err))
		SendFailResponse(ctx, err)
		return
	}
	// 返回结果
	SendSuccessResponse(ctx, nil)
	return
}
