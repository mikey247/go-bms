package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/mikey247/go-bms/pkg/models"
	"github.com/mikey247/go-bms/pkg/utils"
)

var NewBook models.Book

func GetBook(res http.ResponseWriter, req *http.Request) {
	books:= models.GetAllBooks()
	response,_ := json.Marshal(books)

	res.Header().Set("Content-Type","pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func GetBookById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]

	//parsing params from string to int
	ID,err := strconv.ParseInt(bookId,0,0)
	if err!=nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookById(ID)
	response,_ := json.Marshal(bookDetails)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	Book:= &models.Book{}
	// fmt.Println(Book)
	utils.ParseBody(req, Book)
	b:= Book.CreateBook()
	response, _ := json.Marshal(b)
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func UpdateBook(res http.ResponseWriter, req *http.Request){
	var updateBook = &models.Book{}
	fmt.Println(updateBook)
	utils.ParseBody(req,updateBook)

	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails,db := models.GetBookById(ID)

	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	response, _ := json.Marshal(bookDetails)
	res.Header().Set("Content-Type", "pkglication/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

