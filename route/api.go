package route

import (
	"github.com/gin-gonic/gin"
	"messageBot/controller/manage"
	"messageBot/controller/manage/template"
)

func Api(r *gin.Engine) {

	r.Static("/static", "./static")
	r.LoadHTMLGlob("views/**/*")
	apiRouter := r.Group("manage")

	apiRouter.GET("", manage.Index)

	apiRouter.GET("/template", template.Index)
	apiRouter.POST("/template/create", template.Create)
	apiRouter.GET("/template/list", template.GetList)
	apiRouter.GET("/template/remove", template.Remove)

}
