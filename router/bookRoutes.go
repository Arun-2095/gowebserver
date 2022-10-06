package router

import (
	"webserver/controller"

	"github.com/gorilla/mux"
)

func BookRoutes(muxRouter *mux.Router) {

	router := muxRouter.PathPrefix("/book").Subrouter()

	router.HandleFunc("", controller.GetBookList).Methods("GET")
	router.HandleFunc("", controller.CreateBook).Methods("POST")
	router.HandleFunc("/{id}", controller.GetBook).Methods("GET")

}
