package controller

import (
	"encoding/json"
	"net/http"
	"webserver/model"
)

var authQuery = model.AutherQueries{}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {

	var requestBody, existAuthor model.Author

	json.NewDecoder(r.Body).Decode(&requestBody)

	_, result := authQuery.GetAuthor(requestBody.Name)

	if result.Error != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ResponseBuilder("something went wrong", http.StatusUnauthorized, result.Error))
	} else if result.RowsAffected == 1 {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(model.ResponseBuilder("Author already registered", http.StatusConflict, &existAuthor))

	} else {

		_, result := authQuery.CreateAuthor(&requestBody)

		if result.Error != nil {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(model.ResponseBuilder("Failed while creating Record", http.StatusConflict, result.Error))
		} else {

			json.NewEncoder(w).Encode(model.ResponseBuilder("data fetched successfully", http.StatusOK, &requestBody))
		}

	}

}

func GetAuthors(w http.ResponseWriter, r *http.Request) {

	authorsList, queryResult := authQuery.GetAuthors()

	if queryResult.Error != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ResponseBuilder("something went wrong", http.StatusUnauthorized, queryResult.Error))
	} else {

		json.NewEncoder(w).Encode(model.ResponseBuilder("author list fetched successfully", http.StatusOK, &authorsList))

	}
}

func EditAuthor(w http.ResponseWriter, r *http.Request) {

	var requestBody model.Author

	json.NewDecoder(r.Body).Decode(&requestBody)

	_, result := authQuery.EditAuthor(&requestBody)

	if result.Error != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ResponseBuilder("something went wrong", http.StatusUnauthorized, result.Error))
	} else {
		json.NewEncoder(w).Encode(model.ResponseBuilder("author Detail updated successfully", http.StatusOK, "successfully updated"))

	}
}
