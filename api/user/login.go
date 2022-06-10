package api

import (
	"ChineseChess/model"
	service "ChineseChess/service"
	"fmt"
	"github.com/gin-gonic/gin"
	_"github.com/go-sql-driver/mysql"
	"net/http"
)

func Login(c *gin.Context){
	//读取用户参数
	var u model.User
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"message":"出现错误，请联系管理员",
		})
	}
	//判断参数是否合法，用户名最长15位，不为空，密码长于6位
	//判断账号密码是否合法,不为空就行...好吧，密码要大于等于六位，用户名不能大于15位且用户名不能重复
	if (u.Name == "" || u.Password == "") {
		c.JSON(http.StatusOK,gin.H{
			"message":"亲，用户名或密码不能为空哦",
		})
		return
	}
	if len(u.Name) >= 15 {
		c.JSON(http.StatusOK,gin.H{
			"message":"亲，用户名最长为15位哦",
		})
		return
	}
	if len(u.Password) < 6 {
		c.JSON(http.StatusOK,gin.H{
			"message":"亲，密码最短为6位哦",
		})
		return
	}
	//参数合法，判断用户是否存在
	um := service.JudgeUser(u.Name)
	if um.ID == 0 {
		c.JSON(http.StatusOK,gin.H{
			"message":"用户不存在",
		})
		return
	}
	//判断用户是否已经激活
	if um.ActivationCode != "0" {
		c.JSON(http.StatusOK,gin.H{
			"message":"用户未激活",
		})
		return
	}
	//判断用户密码是否正确
	b := service.JudgePw(um.Password,um.Salt,u.Password)
	if !b {
		c.JSON(http.StatusOK,gin.H{
			"message":"密码错误",
		})
		return
	}
	//用户存在且激活返回token
	//在后台打印日志
	fmt.Println(um.Name,"登录成功")
	//返回token
	tokenstring,_ := service.GenerateToken(um.Name)
	c.JSON(http.StatusOK,gin.H{
		"code" : 2000,
		"message" : "欢迎回来，"+um.Name,
		"data" : gin.H{
			"token" : tokenstring,
		},
	})
}
