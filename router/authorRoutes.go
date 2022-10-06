package router

import (
	"webserver/controller"

	"github.com/gorilla/mux"
)

func AuthorRoutes(muxRouter *mux.Router) {

	router := muxRouter.PathPrefix("/author").Subrouter()

	router.HandleFunc("", controller.CreateAuthor).Methods("POST")

}
