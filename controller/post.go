package controller

import (
	"gin-web/logic"
	"gin-web/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePost 创建帖子
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

// PostDetail 帖子详情
func PostDetail(ctx *gin.Context) {

	// 参数获取并校验合法性
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64) // 解析成功则表示是数字字符串
	if err != nil {
		SendFailResponse(ctx, err)
		return
	}

	// 获取id对应帖子
	_ = id
}

// PostList 帖子列表
func PostList(ctx *gin.Context) {

	// 获取分页（ 分页参数获取逻辑可以封装成函数，供多处调用 ）
	page := ctx.Query("page")
	pageSize := ctx.Query("page_size")
	_ = page
	_ = pageSize
}
