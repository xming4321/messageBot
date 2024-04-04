package template

import (
	"github.com/gin-gonic/gin"
	"messageBot/controller"
)

func Index(ctx *gin.Context) {
	data := make(map[string]interface{})
	controller.RenderView(ctx, "template_list.html", data)
}
