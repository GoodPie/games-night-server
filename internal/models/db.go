package models

import (
   "gorm.io/gorm"
   "gorm.io/driver/sqlite"
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