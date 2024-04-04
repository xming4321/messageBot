package template

import (
	"github.com/gin-gonic/gin"
	"messageBot/controller"
	"messageBot/helper"
	service "messageBot/service"
)

type RemoveTemplate struct {
	Id int `form:"id" json:"id" binding:"required,min=1"`
}

func Remove(ctx *gin.Context) {
	removeForm := RemoveTemplate{}

	if err := ctx.ShouldBind(&removeForm); err != nil {
		helper.ErrorfLogger(ctx, "check params error:%+v", err)
		controller.RenderResponse(ctx, -1, err.Error(), nil)
		return
	}

	tplService, err := service.NewTemplateService(ctx)
	if err != nil {
		helper.ErrorfLogger(ctx, "Remove template error:%+v", err)
		controller.RenderResponse(ctx, -1, err.Error(), nil)
		return
	}

	err = tplService.Remove(ctx, removeForm.Id)
	controller.RenderResponseSuccess(ctx, nil)
}
