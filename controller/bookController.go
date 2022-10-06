package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webserver/config"
	"webserver/model"

	"github.com/gorilla/mux"
)

func GetBookList(w http.ResponseWriter, r *http.Request) {
	var dataBase = config.GetDatabase()

	var book []model.Books

	dataBase.Find(&book)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)

}

func GetBook(w http.ResponseWriter, r *http.Request) {

	var dataBase = config.GetDatabase()

	var book model.Books

	params := mux.Vars(r)

	id := params["id"]

	dataBase.Find(&book, id)

	json.NewEncoder(w).Encode(book)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	var book model.Books

	// var dB = config.GetDatabase()

	json.NewDecoder(r.Body).Decode(&book)

	// var dataBase = config.GetDatabase()

	json.NewDecoder(r.Body).Decode(&book)

	fmt.Println(book, "book")

}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	var book model.Books

	var dB = config.GetDatabase()

	json.NewDecoder(r.Body).Decode(&book)

	result := dB.Create(&book)

	json.NewEncoder(w).Encode(result)

}
