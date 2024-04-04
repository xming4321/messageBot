package manage

import (
	"github.com/gin-gonic/gin"
	"messageBot/controller"
	"messageBot/helper"
	"messageBot/service"
)

type GetMessageList struct {
	CurrentPage int `form:"currentPage" json:"currentPage" binding:"omitempty,min=1"`
	PageSize    int `form:"pageSize" json:"pageSize" binding:"omitempty,min=1"`
}

func Index(ctx *gin.Context) {
	getForm := GetMessageList{}

	if err := ctx.ShouldBind(&getForm); err != nil {
		helper.ErrorfLogger(ctx, "check params error:%+v", err)
		controller.RenderResponse(ctx, -1, err.Error(), nil)
		return
	}

	msgService, err := service.NewTgMessageService(ctx)
	if err != nil {
		helper.ErrorfLogger(ctx, "new message service error:%+v", err)
		controller.RenderResponse(ctx, -1, err.Error(), nil)
		return
	}
	currentPage := controller.CurrentPage
	if getForm.CurrentPage != 0 {
		currentPage = getForm.CurrentPage
	}
	pageSize := controller.PageSize
	if getForm.PageSize != 0 {
		pageSize = getForm.PageSize
	}

	total, pageCount, err := msgService.GetPageInfo(ctx, pageSize)
	if err != nil {
		helper.ErrorfLogger(ctx, "GetPageInfo error:%+v", err)
		controller.RenderResponse(ctx, -1, err.Error(), nil)
		return
	}

	dataList, err := msgService.GetList(ctx, currentPage, pageSize)
	if err != nil {
		helper.ErrorfLogger(ctx, "GetList company error:%+v", err)
		controller.RenderResponse(ctx, -1, err.Error(), nil)
		return
	}
	data := make(map[string]interface{})
	data["pageSize"] = pageSize
	data["currentPage"] = currentPage
	data["totalCount"] = total
	data["totalPages"] = pageCount
	data["list"] = dataList
	controller.RenderView(ctx, "index.html", data)
}
