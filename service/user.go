/**
* @Author: gxj
* @Data: 2022/7/17-3:16
* @DESC: Contains user-related business logic,包含了用户相关的业务逻辑
**/

package service

import (
	"ChineseChess/dao"
	"ChineseChess/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//JudgeEmail 判断邮箱格式
func JudgeEmail(email string)bool{
	//邮箱格式的正则表达式，由CSDN搜索得到
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	//转译正则表达式
	reg := regexp.MustCompile(pattern)
	//返回正则表达式匹配结果
	return reg.MatchString(email)
}

//JudgeUsername 判断用户名是否存在
func JudgeUsername(name string)bool{
	//连接数据库
	db := dao.Link()

	//从数据库中提取数据
	u := dao.NameQuery(db,name)
	if u.ID == 0 {
		return true
	}
	defer db.Close()
	return false

}

//发送邮箱激活码
func SendActivationCode(email string)string{
	//绑定邮箱的地址,发送激活码的邮箱
	senderEmail := "323150736@qq.com"
	//设置smtp，qq邮箱地址及端口，授权码
	smt := "smtp.qq.com"
	smtpPort := ":587"
	authorizePassword := "nlfkdkycxypccabg"
	//头部信息
	auth := smtp.PlainAuth("", senderEmail, authorizePassword,smt)
	//设置邮件发送内容类型
	contentType := "Content-Type: text/plain;charset=UTF-8"
	//转变收件人邮箱格式为切片
	receiver := []string{email}
	//设置发送信息，发件人，标题，内容
	//发件人
	senderName := "redRockChineseChess-project"
	//标题
	title := "您的激活码"
	//内容，随机四位数的激活码，利用时间戳和rand随机数
	rand.Seed(time.Now().UnixNano())
	var activationCode string
	for i:=0;i<4;i++{
		ve := rand.Intn(10)
		activationCode = activationCode +strconv.Itoa(ve)

	}
	//配置发送信息格式
	msg := []byte("To: " + strings.Join(receiver, ",") + "\r\nFrom: " + senderName +
		"<" + senderEmail + ">\r\nSubject: " + title + "\r\n" + contentType + "\r\n\r\n" + activationCode)
	//调用接口发送信息,并处理错误
	err := smtp.SendMail(smt+smtpPort,auth, senderEmail,receiver,msg)
	if err != nil {
		fmt.Println(err.Error(),"  ",time.Now().Format("2006-01-02 15:04:05"))
		return " "
	}
	return activationCode
}

//保存用户
func SaveNewUser(u model.User)bool{
	//对时间进行赋值
	u.CreateTime = time.Now().Add(10*time.Minute)
	//连接数据库
	db := dao.Link()

	//将数据存入数据库中
	err := dao.SaveNewUser(db,u)
	defer db.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

//获取激活码
func ObtainCode (n string)model.User{
	//连接到数据库
	db := dao.Link()

	//通过name查找数据库中信息
	u :=dao.NameQuery(db,n)
	defer db.Close()
	return u
}


//更新激活后的用户信息，并判断是否过期
func UpdateNewUser(u model.User)error{
	//连接数据库
	db := dao.Link()

	//更新数据
	err := dao.UpdateNewUser(db,u)
	defer db.Close()
	return err
}

//使用md5加密用户密码，密码+盐值
func EncryPs(ps string)string{
	h := md5.New()
	h.Write([]byte(ps))
	return hex.EncodeToString(h.Sum(nil))
}

func JudgePw(pw,salt,in string)bool{
	h := md5.New()
	h.Write([]byte(in+salt))
	strm := hex.EncodeToString(h.Sum(nil))
	if strm == pw {
		return true
	}
	return false
}


//判断用户是否存在，读取数据库中信息
func JudgeUser(n string)model.User{
	//连接到数据库
	db := dao.Link()

	//通过用户名，读取信息
	u := dao.NameQuery(db,n)
	defer db.Close()
	return u
}




