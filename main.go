package main

import (
	"fmt"
	"log"
	"net/http"
	"webserver/config"
	"webserver/model"
	"webserver/router"

	"github.com/gorilla/mux"
)

func main() {

	mainRouter := mux.NewRouter()

	router.BookRoutes(mainRouter)

	router.AuthorRoutes(mainRouter)

	// middleware , quering database;

	databaseInit()

	log.Fatal(http.ListenAndServe(":8080", mainRouter))

}

func databaseInit() {
	fmt.Println("INITIATING DATABASE CONNECTION")
	config.Connect()
	database := config.GetDatabase()
	database.AutoMigrate(&model.Books{}, &model.Author{})
	fmt.Println("DATABASE CONNECTION INITIATED")

}
