package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)


type User struct {
	ID int `gorm:"primary_key"`
	Name string `gorm:"unique"`
	Password string
	Email string
	ActivationCode string
	CreateTime time.Time
	Salt string
}

//JWT相关结构体
//声明claims结构体,自定义字段添加用户名
type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}
