package dao

import (
	"ChineseChess/model"
	"github.com/jinzhu/gorm"
)

func NameQuery(db *gorm.DB,n string)model.User{
	var u model.User
	db.Where(&model.User{Name: n}).First(&u)
	return u

}

func SaveNewUser(db *gorm.DB,u model.User)error{
	err := db.Create(&u).Error
	return err
}
