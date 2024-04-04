package main

import (
	"github.com/gin-gonic/gin"
	"messageBot/conf"
	"messageBot/helper"
	"messageBot/route"
	"messageBot/service"
)

func main() {
	r := gin.Default()
	helper.InitLog()
	helper.InitMysql()
	helper.InitTgBot()
	service.StartReceiveTgMessage()
	route.Api(r)
	err := r.Run(":" + conf.Conf.Port)
	if err != nil {
		helper.PanicfLogger(nil, "run error %+v", err)
	}
}
