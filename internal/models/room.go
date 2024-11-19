package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	SessionId uint
	Code string `gorm:"unique"`
	Name string
	// Conscious decision to only add points to first three places
	FirstPoints int  
	SecondPoints int
	ThirdPoints int
	Users []User `gorm:"many2many:room_users"`
}