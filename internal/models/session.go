package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	CreatorId uint
	Rooms []Room
	Users []User `gorm:"many2many:session_users"`
}
