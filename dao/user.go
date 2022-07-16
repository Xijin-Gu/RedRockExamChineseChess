/**
* @Author: gxj
* @Data: 2022/7/17-3:22
* @DESC: Contains modifications to user data,包含了对用户类数据的修改
**/

package dao

import (
	"ChineseChess/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

//NameQuery 通过用户名字查找用户相关信息
func NameQuery(db *gorm.DB,n string)model.User{
	var u model.User
	db.Where(&model.User{Name:n}).First(&u)
	return u

}

//SaveNewUser 保存新用户信息
func SaveNewUser(db *gorm.DB,u model.User)error{
	err := db.Create(&u).Error
	return err
}

//UpdateNewUser 更新用户信息
func UpdateNewUser(db *gorm.DB,u model.User)error{
	fmt.Println("uid",u.ID)
	err := db.Model(&model.User{}).Where("id = ?",u.ID).Update(&u).Error
	return err
}
