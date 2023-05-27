package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"book-management-api.io/packages/models"
	"book-management-api.io/packages/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	NewBook := &models.Book{}       // a model of type book
	utils.ParseBody(r, NewBook)     // parse book details from url
	b := models.CreateBook(NewBook) // create query forcreating book into db

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	NewBooks := models.GetAllBooks() //calling Mysql function to get books
	res, _ := json.Marshal(NewBooks) // encoding data to json for sending response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookById(w http.ResponseWriter, r *http.Request) {
	//First parse the ID of book from URl and save in ID
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing")
	}

	//Then send the ID to models file to GetBooksbyId to get book details from DB
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	//Find that bOok in Db , update required fields
	UpdateBook := &models.Book{} // a model of type book
	utils.ParseBody(r, UpdateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing")
	}
	bookDetails, db := models.GetBookById(ID) //fetching book from db bu id
	if UpdateBook.Name != "" {
		bookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		bookDetails.Publication = UpdateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing")
	}

	bookDetails := models.DeleteBook(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
