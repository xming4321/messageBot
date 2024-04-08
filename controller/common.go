package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义返回
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// 错误定义
var ErrorResponseNotExist = Response{Status: 404, Message: "记录不存在"}

// 常量定义
const PageSize = 20
const CurrentPage = 1

func RenderResponse(ctx *gin.Context, status int, msg string, data interface{}) {
	ctx.JSON(200, Response{
		Status:  status,
		Message: msg,
		Data:    data,
	})
	return
}

func RenderError(ctx *gin.Context, res Response) {
	ctx.JSON(200, res)
}

func RenderResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, Response{
		Status:  200,
		Message: "成功",
		Data:    data,
	})
	return
}

func RenderView(ctx *gin.Context, tpl string, data interface{}) {
	ctx.HTML(http.StatusOK, tpl, data)
}
