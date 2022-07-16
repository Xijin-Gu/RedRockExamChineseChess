/**
* @Author: gxj
* @Data: 2022/7/17-3:24
* @DESC: Contains game-related structures,包含了游戏相关的结构体
**/

package model

type Game struct {
	ID int `gorm:"primary_key"`//本局游戏ID
	HouseID int				//房间ID
	Mover string			//游戏当前移动方
	Waiter       string		//游戏当前非移动方
	WinLose      string		//游戏胜负状况
	Checkerboard string		//游戏棋盘信息
}

