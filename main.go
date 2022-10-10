package main

import (
	"fmt"
	"log"
	"net/http"
	"webserver/config"
	"webserver/middleware"
	"webserver/model"
	"webserver/router"

	"github.com/gorilla/mux"
	"gorm.io/gorm/logger"
)

func main() {

	mainRouter := mux.NewRouter()

	mainRouter.Use(middleware.LoggingMiddleware, middleware.SettingDefaultHeader)

	router.BookRoutes(mainRouter)

	router.AuthorRoutes(mainRouter)

	//TODO: validate requestbody, authtoken, relation db queries;

	databaseInit()

	log.Fatal(http.ListenAndServe(":8080", mainRouter))

}

func databaseInit() {
	fmt.Println("INITIATING DATABASE CONNECTION")
	config.Connect()
	database := config.GetDatabase()

	database.AutoMigrate(&model.Books{}, &model.Author{})
	fmt.Println("DATABASE CONNECTION INITIATED")
	database.Logger.LogMode(logger.Info)

}
