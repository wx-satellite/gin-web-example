package controller

import (
	"gin-web/pkg/translator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

const ContextUserIdKey = "userId"

/**
{
	"code": 1000, // 错误码
	"msg": xxx, // 描述
	"data": {} // 数据
}
*/

/**
错误的处理方式：
1. 维护一份 Code 和 Message 的 map
2. 维护自定义错误类型：
	type Error struct {
		Code    int
		Message string //可以展示给前端的信息
		Err     string `json:"-"` //具体错误信息，可能包含敏感数据
		Caller  string `json:"-"` //调用栈
	}
*/
type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func successResponse(message string, data interface{}) Response {
	return Response{Message: message, Data: data}
}

func SendSuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, successResponse("成功", data))
}

func SendFailResponse(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, failResponse(err))
	return
}

func failResponse(err error) Response {
	var message string
	switch e := err.(type) {
	// validator 包错误信息翻译
	case validator.ValidationErrors:
		messages := e.Translate(translator.GetTranslator())
		for _, value := range messages {
			message += value + "，"
		}
	default:
		message = err.Error()
	}
	return Response{Message: strings.Trim(message, "，")}
}
