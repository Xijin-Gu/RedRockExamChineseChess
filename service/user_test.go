/**
* @Author: gxj
* @Data: 2022/7/18-14:12
* @DESC:
**/

package service

import (
	"ChineseChess/model"
	"crypto/md5"
	"encoding/hex"
	"github.com/go-playground/assert/v2"
	"testing"
)

//TestJudgeEmail 测试邮箱格式
func TestJudgeEmail(t *testing.T) {
	email := []struct{
		account string
		is      bool
	}{
		{"1273447417@qq.com",true},
		{"323150736@qq.com",true},
		{"2021210577",false},
	}

	for _,v := range email{
		t.Run(v.account, func(t *testing.T) {
			got := JudgeEmail(v.account)
			assert.Equal(t,got,v.is)
		})
	}
}


func TestJudgeUsername(t *testing.T) {
	name := "guxijin"
	got := JudgeUser(name)
	assert.Equal(t,got.ID,3)
}

//
//func TestSendActivationCode(t *testing.T) {
//	codes := []struct{
//		email string
//		re	string     //返回数值
//	}{
//		{"1273447417@qq.com","[0-9]"},
//		{"323150736"," "},
//	}
//	for _,v := range codes{
//		t.Run(v.email, func(t *testing.T) {
//			got := SendActivationCode(v.email)
//			assert.Equal(t,got,v.re)
//		})
//	}
//
//}


func TestSaveNewUser(t *testing.T) {
	user := []struct{
		user model.User
		is bool
	}{
		{model.User{Name: "guxijin"},false},
		{model.User{},false},
	}
	for _,v := range user{
		t.Run(v.user.Name, func(t *testing.T) {
			got := SaveNewUser(v.user)
			assert.Equal(t,got,v.is)
		})
	}
}

func TestObtainCode(t *testing.T) {
	name := "guxjin"
	got := ObtainCode(name)
	assert.Equal(t,got,model.User{})
}

func TestUpdateNewUser(t *testing.T) {
	g := model.User{Name: "guxijin"}
	got := UpdateNewUser(g)
	assert.Equal(t,got,nil)
}

func TestEncryPs(t *testing.T) {
	ps := "123456"
	h := md5.New()
	h.Write([]byte(ps))
	want := hex.EncodeToString(h.Sum(nil))
	got := EncryPs(ps)
	assert.Equal(t,got,want)
}

