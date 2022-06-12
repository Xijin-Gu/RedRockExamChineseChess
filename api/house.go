package api

import (
	"ChineseChess/model"
	"ChineseChess/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//创建房间
func CreateHouse(c *gin.Context){
	//验证token
	name := c.MustGet("name").(string)
	//token有效，则提取token中的用户名
	if name == "" {
		c.JSON(http.StatusOK,gin.H{
			"message" : "请登录",
		})
		return
	}
	//创建房间
	var h model.House
	h.CreateName = name
	//存储房间信息，返回房间ID
	id := service.CreateHouse(h)
	c.JSON(http.StatusOK,gin.H{
		"message":"您的房间id为:"+strconv.Itoa(id),
	})
}
func JoinHouse(c *gin.Context){
	//验证token
	name := c.MustGet("name").(string)
	//token有效，则提取token中的用户名
	if name == "" {
		c.JSON(http.StatusOK,gin.H{
			"message" : "请登录",
		})
		return
	}
	houseID,_ := strconv.Atoi(c.Param("id"))

	//判断房间是否存在
	b := service.JudgeHouse(houseID)
	if !b {
		c.JSON(http.StatusOK,gin.H{
			"message":"房间不存在",
		})
		return
	}
	//判断是否为房主
	bo := service.JudgeHouseOwner(name,houseID)
	if bo {
		c.JSON(http.StatusOK,gin.H{
			"message":"欢迎房主",
		})
	} else {
		//录入信息
		err := service.SaveParticipate(name,houseID)
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"message":"出现错误,请联系管理员",
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"message":"欢迎加入,"+name,
		})
	}

	//录入状态信息
	state := c.Query("state")
	if (state == "0" || state == "1") {
		err := service.SaveState(state,bo,houseID)
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"message":"出现错误，请联系管理员",
			})
			return
		}
	}

	//判断两人的状态
	bs := service.JudgeState(houseID)
	if !bs {
		c.JSON(http.StatusOK,gin.H{
			"message":"等待准备中",
		})
		return
	}
	//创建一个游戏
	gameID := service.InitGameAndPiece(houseID)
	c.JSON(http.StatusOK,gin.H{
		"message":"游戏id为:"+strconv.Itoa(gameID),
	})

}



