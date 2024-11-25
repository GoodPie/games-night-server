package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("games.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	return DB.AutoMigrate(&User{}, &Session{}, &Room{})
}
