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
		userGroup.POST("/regist/sendactivation",api.SendActivationCode)	//注册界面，会发送邮箱激活码，需要用户使用用户名和激活码进行激活
		userGroup.POST("/regist/active",api.VerifyActivationCode)		//激活界面，用户提供用户名和激活码。进行账号激活
		userGroup.POST("/login",api.Login)								//登录接口，用户进行登录
	}


	//定义房间组
	userGroupH := engine.Group("/house")
	{
		userGroupH.POST("/create",service.VerifyJWT(),api.CreateHouse)	//创建房间，用户创建一个新的游戏房间
		userGroupH.GET("/:id",service.VerifyJWT(),api.JoinHouse)			//加入房间，用户可以通过房间id加入房间
	}


	//定义游戏界面,协议升级为websocket
	engine.POST("/game/:id",service.VerifyJWT(),api.ChessGame)		//游戏界面，供用户游玩。

	//启动
	engine.Run(":8090")




}
