package dao

import (
	"ChineseChess/model"
	"github.com/jinzhu/gorm"
)

func QueryGameDataByID(db *gorm.DB,gameID int)model.Game{
	var gameData model.Game
	db.Model(&model.Game{}).Where(&model.Game{ID: gameID}).First(&gameData)
	return gameData

}

func UpdateGameData(db *gorm.DB,game model.Game){
	db.Model(&model.Game{}).Where(&model.Game{ID: game.ID}).Update(&game)
}
