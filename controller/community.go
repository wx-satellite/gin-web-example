package controller

import (
	"fmt"
	"gin-web/logic"
	"gin-web/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// ----- 社区相关 -----

// Communities 社区列表
// @Description 社区列表接口
// @Tags 社区相关接口
// @Param Authorization header string true "Bearer JWT"
// @Param object query request.CommunityList false "分页结构体"
// @Success 200 {object} controller.Response "返回值"
// @Failure 200 {object} controller.Response "返回值"
// @Router /community [get]
func Communities(ctx *gin.Context) {
	// 参数验证
	in := &request.CommunityList{
		//Page:     1,
		//PageSize: 20,
	}
	if err := ctx.ShouldBind(in); err != nil {
		SendFailResponse(ctx, err)
		return
	}
	fmt.Println(in.Page)
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
