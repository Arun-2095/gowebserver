package main

import (
	"log"
	"net/http"
	"webserver/controller"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", controller.HomeController).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}
