package dao

import (
	"ChineseChess/model"
	"github.com/jinzhu/gorm"
)

func CreateHouse(db *gorm.DB,n string)int{
	db.Where(&model.House{CreateName: n}).Create(&model.House{CreateName: n})
	var h model.House
	db.Where(&model.House{CreateName: n}).First(&h)
	return h.ID
}


func QueryHouse(db *gorm.DB,id int)(error,model.House){
	var h model.House
	err := db.Where(&model.House{ID: id}).First(&h).Error
	return err,h
}

func UpdateHouse(db *gorm.DB,house model.House)error{
	err := db.Model(&model.House{}).Where(&model.House{ID: house.ID}).Update(&house).Error
	return err
}