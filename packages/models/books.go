package models

import (
	"book-management-api.io/packages/configs"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	configs.ConnectToDb() //Calling function from config
	db = configs.GetDB()
	db.AutoMigrate(&Book{}) // Send data to DB
}
