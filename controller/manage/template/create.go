package template

import (
	"github.com/gin-gonic/gin"
	"messageBot/controller"
	"messageBot/helper"
	"messageBot/model"
	service2 "messageBot/service"
	"time"
)

type CreateTemplate struct {
	Command string `form:"command" json:"command" binding:"omitempty"`
	Reg     string `form:"name" json:"reg" binding:"omitempty"`
	Reply   string `form:"reply" json:"reply" binding:"required,min=1"`
}

func Create(ctx *gin.Context) {
	createForm := CreateTemplate{}

	if err := ctx.ShouldBind(&createForm); err != nil {
		helper.ErrorfLogger(ctx, "check params error:%+v", err)
		controller.RenderResponse(ctx, -1, err.Error(), nil)
		return
	}

	if createForm.Command == "" && createForm.Reg == "" {
		controller.RenderResponse(ctx, -1, "command or reg must choose one!", nil)
		return
	}

	company := model.Template{
		Command:    createForm.Command,
		Reg:        createForm.Reg,
		Reply:      createForm.Reply,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	service, err := service2.NewTemplateService(ctx)
	if err != nil {
		helper.ErrorfLogger(ctx, "new template service error:%+v", err)
		controller.RenderResponse(ctx, -1, err.Error(), nil)
		return
	}

	_, err = service.Create(ctx, company)
	if err != nil {
		helper.ErrorfLogger(ctx, "company create error:%+v", err)
		controller.RenderResponse(ctx, -1, err.Error(), nil)
		return
	}
	controller.RenderResponseSuccess(ctx, nil)
}
