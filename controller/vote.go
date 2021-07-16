package controller

import (
	"gin-web/request"
	"github.com/gin-gonic/gin"
)

func VotePost(ctx *gin.Context) {
	in := new(request.VotePostRequest)
	if err := ctx.ShouldBindJSON(in); err != nil {
		SendFailResponse(ctx, err)
		return
	}
	// 填充当前登陆的用户ID
	in.UserId = getCurrentUserId(ctx)

}
