package service

import (
	"ChineseChess/dao"
	"ChineseChess/model"
	_ "github.com/go-sql-driver/mysql"
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




func InitGameAndPiece(id int)int{
	//连接数据库
	db := dao.Link()
	//通过id读取House表信息
	_,house := dao.QueryHouse(db,id)
	//开始初始化信息
	//初始化游戏表
	var gameData model.Game
	gameData.HouseID = house.ID
	gameData.Mover = house.CreateName
	gameData.Waiter = house.ParticipateName
	gameData.Winlose = "无"
	gameData.Checkerboard = "151413121112131415\n000000000000000000\n001600000000001600\n170017001700170017\n000000000000000000\n000000000000000000\n270027002700270027\n002600000000002600\n000000000000000000\n252423222122232425"
	gameDataMysql := dao.SaveNewGame(db,gameData)
	return gameDataMysql.ID
}