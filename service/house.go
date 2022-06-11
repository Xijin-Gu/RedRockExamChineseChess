package service

import (
	"ChineseChess/dao"
	"ChineseChess/model"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
)

//根据用户名创建一个房间
func CreateHouse(h model.House)int{
	//连接数据库
	db := dao.Link()
	//创建一个新房间词条
	id := dao.CreateHouse(db,h.CreateName)
	return id
}

//判断房间是否存在
func JudgeHouse(id int)bool{
	//连接数据库
	db := dao.Link()
	//查询数据库
	err,_ := dao.QueryHouse(db,id)
	if err != nil {
		if err.Error() == "record not found" {
			return false
		}
	}
	return true
}



//判断是否为房主
func JudgeHouseOwner(n string,id int)bool{
	//连接数据库
	db := dao.Link()
	//查询数据库
	_,house := dao.QueryHouse(db,id)
	if n == house.CreateName {
		return true
	}
	return false
}

//录入参与者信息
func SaveParticipate(n string,id int)error{
	//连接数据库
	db := dao.Link()
	//查询数据库
	_,house := dao.QueryHouse(db,id)
	//更新信息
	house.ParticipateName = n
	//更新数据库
	err := dao.UpdateHouse(db,house)
	return err

}

//录入状态信息
func SaveState(st string,b bool,id int)error{
	//连接数据库
	db := dao.Link()
	//查询数据库
	_,house := dao.QueryHouse(db,id)
	//更新信息
	if b {
		house.CreateState = st
	} else {
		house.ParticipateState = st
	}
	//更新数据库
	err := dao.UpdateHouse(db,house)
	return err
}

//判断两人是否准备完毕
func JudgeState(id int)bool{
	//连接数据库
	db := dao.Link()
	//查询数据库
	_,house := dao.QueryHouse(db,id)
	//判断状态信息
	if (house.CreateState == "1" && house.ParticipateState == "1") {
		return true
	}
	return false
}


func CreateGame(id int)int{
	//连接数据库
	//db := dao.Link()
	//创建新的游戏
	return rand.Int()
}
