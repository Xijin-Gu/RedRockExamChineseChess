package model

import "time"

type User struct {
	ID int `gorm:"primary_key"`
	Name string `gorm:"unique"`
	Password string
	Email string
	ActivationCode string
	CreateTime time.Time
	Salt string
}
