/**
* @Author: gxj
* @Data: 2022/7/17-3:25
* @DESC: Contains house-related structures,包含了房间相关的结构体
**/

package model

type House struct {
	ID int `gorm:"primary_key"`	//房间ID
	CreateName string			//房间创建者
	ParticipateName string		//房间参与者
	CreateState string			//创建者是否已准备
	ParticipateState string		//参与者是否已准备
}
