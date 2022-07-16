/**
* @Author: gxj
* @Data: 2022/7/17-3:23
* @DESC: Defines the function to connect to the database,定义了连接数据库的函数
**/

package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

//Link 连接到数据库
func Link() *gorm.DB {
	db,err := gorm.Open("mysql","root:2002@/chinesechess?charset=utf8&loc=Local&parseTime=True")
	if err != nil {
		fmt.Println("err:",err)
	}
	return db
}
