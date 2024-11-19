package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirebaseId string `gorm:"unique"`
	Name       string
	IsAdmin	bool
	IsGuest	bool
    Sessions  []Session `gorm:"many2many:user_sessions"`
}
