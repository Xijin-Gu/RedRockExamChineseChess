package model

type House struct {
	ID int `gorm:"primary_key"`
	CreateName string
	ParticipateName string
	CreateState string
	ParticipateState string
}
