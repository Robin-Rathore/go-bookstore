package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/Robin-Rathore/go-bookstore/pkg/utils"
	"github.com/Robin-Rathore/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json") // ✅ Fixed typo
	w.WriteHeader(http.StatusOK) // ✅ Fixed StatusOK
	w.Write(res) // ✅ Fixed function call
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json") // ✅ Fixed typo
	w.WriteHeader(http.StatusOK) // ✅ Fixed StatusOK
	w.Write(res) // ✅ Fixed function call
}

func CreateBook(w http.ResponseWriter, r *http.Request) { // ✅ Fixed incorrect Request type
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK) // ✅ Fixed StatusOK
	w.Write(res) // ✅ Fixed function call
}
