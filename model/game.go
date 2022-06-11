package model

type Game struct {
	ID int `gorm:"primary_key"`
	HouseID int
	Mover string
	Waiter string
	Winlose string
	Checkerboard string
}

