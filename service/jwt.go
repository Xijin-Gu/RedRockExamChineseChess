package service

import (
"ChineseChess/model"
"errors"
"fmt"
"github.com/dgrijalva/jwt-go"
"github.com/gin-gonic/gin"
"net/http"
"strings"
"time"
)

//定义Jwt生效时间和secret
var JwtEffectiveTime = time.Hour*48
var Secret =[]byte("RedRockChineseChess")

//jwt生成
func GenerateToken(name string)(string,error){
	//创建声明
	c := model.Claims{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(JwtEffectiveTime).Unix(),
			Issuer: "RedRockChineseChess-project",
		},
	}
	fmt.Println("claims",c)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	return token.SignedString(Secret)

}


//JWT解析
func ParseJWT(tokenstring string)(*model.Claims,error){
	//解析token
	token,err := jwt.ParseWithClaims(tokenstring,&model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret,nil
	})
	//处理错误
	if err != nil {
		return nil, err
	}
	//验证是否token有效
	if claims,ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claims,nil
	}
	return nil, errors.New("token无效")
}


//JWT认证
func VerifyJWT()func(c *gin.Context){
	return func(c *gin.Context) {

		//获取含有token信息的头部Authorazition部分
		authorization := c.Request.Header.Get("Authorization")
		fmt.Println(authorization)
		if authorization == "" {
			c.JSON(http.StatusOK,gin.H{
				"code" : 2003,
				"message" : "Authorazition为空",
			})
			c.Abort()
			return
		}

		//提取token信息段
		JwtInformation := strings.SplitN(authorization," ",2)
		//验证auth信息段是否合法
		if !(len(JwtInformation) == 2 && JwtInformation[0] == "Bearer") {
			c.JSON(http.StatusOK,gin.H{
				"code" : 2004,
				"message" : "Authorazition格式错误",
			})
			c.Abort()
			return
		}
		//验证token是否有效
		claim,err := ParseJWT(JwtInformation[1])
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"code" : 2005,
				"message" : "token无效",
			})
			c.Abort()
			return
		}
		//将claim信息保存到上下文
		c.Set("name",claim.Name)
		c.Next()
	}
}
