package main

import (
	api "ChineseChess/api/user"
	"ChineseChess/service"
	"github.com/gin-gonic/gin"
)

func main(){
	engine := gin.Default()

	//解决跨域问题
	engine.Use(service.CrossDomain())


	//定义用户组
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/regist/sendactivation",api.SendActivationCode)
	}


	//启动
	engine.Run(":8090")




}
