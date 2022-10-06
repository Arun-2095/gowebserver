package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webserver/config"
	"webserver/model"
)

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var requestBody, existAuthor model.Author

	db := config.GetDatabase()

	json.NewDecoder(r.Body).Decode(&requestBody)

	result := db.Where("name=?", requestBody.Name).Find(&existAuthor)

	if result.Error != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ResponseBuilder("something went wrong", http.StatusUnauthorized, result.Error))
	} else if result.RowsAffected == 1 {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(model.ResponseBuilder("Author already registered", http.StatusConflict, &existAuthor))

	} else {

		result := db.Create(&requestBody)

		if result.Error != nil {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(model.ResponseBuilder("Failed while creating Record", http.StatusConflict, result.Error))
		} else {

			json.NewEncoder(w).Encode(model.ResponseBuilder("data fetched successfully", http.StatusOK, &requestBody))
		}

	}

	fmt.Println(requestBody, result.RowsAffected, result.Error, "create api")

}
