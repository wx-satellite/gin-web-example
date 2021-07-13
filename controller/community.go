package controller

import (
	"gin-web/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// ----- 社区相关 -----

// Communities 社区列表
func Communities(ctx *gin.Context) {
	// 参数验证

	// 业务逻辑处理
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList handle fail", zap.Error(err))
		SendFailResponse(ctx, err)
		return
	}

	// 返回结果
	SendSuccessResponse(ctx, data)
}

// CommunityDetail 社区详情
func CommunityDetail(ctx *gin.Context) {

	// 参数验证
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		SendFailResponse(ctx, err)
		return
	}

	// 业务处理
	obj, err := logic.CommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.CommunityDetail handle fail", zap.Error(err))
		SendFailResponse(ctx, err)
		return
	}

	// 返回结果
	SendSuccessResponse(ctx, obj)
}
