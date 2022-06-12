package api

import (
	"ChineseChess/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

var upgrader = websocket.Upgrader{
	// 这个是校验请求来源,直接return true
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


func ChessGame(c *gin.Context){
	//升级协议
	client, _ := upgrader.Upgrade(c.Writer, c.Request, nil)
	//进入循环
	for {
		//读取用户信息
		gameId := c.Param("id")
		name := c.Query("name")
		start := c.Query("start")
		end := c.Query("end")
		gameID,_ := strconv.Atoi(gameId)
		//判断用户是否合法,并读取游戏表中的信息
		boUser,gameData  := service.JudgeGameUser(name,gameID)
		if !boUser {
			var a = "这不是您加入的游戏"
			err := client.WriteMessage(1,[]byte(a))
			if err != nil {
				log.Println(err)
			}
			return
		}
		//解析棋盘
		var checkerboard [10][9]string
		checkerboard = service.ParseCheckboard(gameData.Checkerboard)
		//走棋后

		gameUpdate := service.MoveChess(start,end,checkerboard,gameData)

				a := service.UpdateChessData(gameUpdate)
				//把a塞进通道进行播报
				err := client.WriteMessage(1,[]byte(a))
				if err != nil {
				log.Println(err)
				}
		time.Sleep(time.Second*10)
	}

}