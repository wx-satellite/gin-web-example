package controller

import (
	"gin-web/pkg/translator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type Response struct {
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
