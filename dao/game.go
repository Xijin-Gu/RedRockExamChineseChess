/**
* @Author: gxj
* @Data: 2022/7/17-3:21
* @DESC: Contains modifications to game data,包含了对游戏类数据的修改
**/

package dao

import (
	"ChineseChess/model"
	"github.com/jinzhu/gorm"
)

//QueryGameDataByID 通过游戏ID查找游戏相关信息
func QueryGameDataByID(db *gorm.DB,gameID int)model.Game{
	var gameData model.Game
	db.Model(&model.Game{}).Where(&model.Game{ID: gameID}).First(&gameData)
	return gameData

}
//UpdateGameData 更新游戏信息
func UpdateGameData(db *gorm.DB,game model.Game){
	db.Model(&model.Game{}).Where(&model.Game{ID: game.ID}).Update(&game)
}
