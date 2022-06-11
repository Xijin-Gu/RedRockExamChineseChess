package main

import (
	"ChineseChess/api"
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
		userGroup.POST("/regist/active",api.VerifyActivationCode)
		userGroup.POST("/login",api.Login)
	}


	//定义房间组
	userGroupH := engine.Group("/house")
	{
		userGroupH.POST("/create",service.VerifyJWT(),api.CreateHouse)
		userGroupH.GET("/:id",service.VerifyJWT(),api.JoinHouse)
	}

	//启动
	engine.Run(":8090")




}
