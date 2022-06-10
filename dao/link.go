package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
)


func Link() *gorm.DB {
	db,err := gorm.Open("mysql","root:2002@/chinesechess?charset=utf8&loc=Local&parseTime=True")
	if err != nil {
		fmt.Println("err:",err)
	}
	return db
}
