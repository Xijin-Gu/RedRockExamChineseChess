/**
* @Author: gxj
* @Data: 2022/7/17-3:22
* @DESC: Contains modifications to house data,包含了对房间类数据的修改
**/

package dao

import (
	"ChineseChess/model"
	"github.com/jinzhu/gorm"
)

//CreateHouse 创建房间数据
func CreateHouse(db *gorm.DB,n string)int{
	db.Where(&model.House{CreateName: n}).Create(&model.House{CreateName: n})
	var h model.House
	db.Where(&model.House{CreateName: n}).First(&h)
	return h.ID
}

//QueryHouse 查询房间信息
func QueryHouse(db *gorm.DB,id int)(error,model.House){
	var h model.House
	err := db.Where(&model.House{ID: id}).First(&h).Error
	return err,h
}

//UpdateHouse 更新房间信息
func UpdateHouse(db *gorm.DB,house model.House)error{
	err := db.Model(&model.House{}).Where(&model.House{ID: house.ID}).Update(&house).Error
	return err
}

//SaveNewGame 第一次存储游戏表
func SaveNewGame(db *gorm.DB,game model.Game)model.Game{
	db.Model(&model.Game{}).Create(&game)
	var gr model.Game
	db.Model(&model.Game{}).Where(&game).First(&gr)
	return gr
}
