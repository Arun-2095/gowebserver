package controller

import (
	"encoding/json"
	"net/http"
	"webserver/model"
)

var books = []model.Book{{Name: "3 idiots"}, {Name: "2 states"}, {Name: "half girlfriend"}}

func HomeController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}
