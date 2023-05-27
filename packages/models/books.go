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
	db.AutoMigrate(&Book{}) //  auto migration for given models,
}

//Functions to talk to database

func CreateBook(b *Book) *Book {
	db.NewRecord(b) //Prepare SQL Query : Done by Gorm Package
	db.Create(&b)   //Execute the query
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books) //Find all records of type Book
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
