package api

import (
	"ChineseChess/model"
	service "ChineseChess/service/user"
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
)

//该go文件共有两个接口，一个接口用来发送激活码，一个接口用来验证激活码


func SendActivationCode(c *gin.Context){
	//获取用户发送信息
	//绑定参数并处理错误
	var u model.User
	err := c.ShouldBind(&u)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK,gin.H{
			"message":"注册出错，请联系管理员",
		})
		return
	}
	//判断参数是否合法
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
	//判断邮箱是否合法
	if !(service.JudgeEmail(u.Email)) {
		c.JSON(http.StatusOK,gin.H{
			"message":"邮箱格式不对哟",
		})
		return
	}
	//判断用户名是否存在
	if !(service.JudgeUsername(u.Name)){
		c.JSON(http.StatusOK,gin.H{
			"message":"用户名已存在请重新输入",
		})
		return
	}
	//发送邮箱激活码
	code := service.SendActivationCode(u.Email)
	//判断激活码发送是否错误
	if code == " " {
		c.JSON(http.StatusOK,gin.H{
			"message":"验证码发送错误请联系管理员",
		})
		return
	}
	u.ActivationCode = code
	//将激活码码存储到数据表中，等待用户激活
	b := service.SaveNewUser(u)
	//判断是否存储成功
	if !b {
		c.JSON(http.StatusOK,gin.H{
			"message":"出现错误，请联系管理员",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"ActivationCode" : code,
		"message":"请尽快激活账号",
	})
}


func VerifyActivationCode(c *gin.Context){
	//获取用户发送信息

	//从数据库中提取信息


	//判断激活码是否正确


	//激活码正确，更新用户表


}