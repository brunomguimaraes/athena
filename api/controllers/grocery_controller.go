package controllers

import (
	"athena/api/models"
	"athena/api/responses"
	validationErrors "athena/api/utils/errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (server *Server) CreateGrocery(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	grocery := models.Grocery{}
	err = json.Unmarshal(body, &grocery)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	grocery.Setup()
	err = grocery.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	groceryCreated, err := grocery.SaveGrocery(server.DB)
	if err != nil {
		formattedError := validationErrors.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, groceryCreated.Id))
	responses.JSON(w, http.StatusCreated, groceryCreated)
}
