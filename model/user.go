/**
* @Author: gxj
* @Data: 2022/7/17-3:25
* @DESC: Contains user-related structures,包含了用户相关的结构体
**/

package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)


type User struct {
	ID int `gorm:"primary_key" json:"id"`	//用户ID
	Name string `gorm:"unique" json:"name"`	//用户名
	Password string			`json:"password"`	//用户密码，加盐后存储
	Email string				`json:"email"`//用户邮箱
	ActivationCode string		//用户验证码，激活后为0
	CreateTime time.Time		//创建时间
	Salt string					//用户密码的盐值
}

//Claims JWT相关结构体
//声明claims结构体,自定义字段添加用户名
type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}
