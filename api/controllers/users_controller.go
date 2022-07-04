package controllers

import (
	"athena/api/models"
	"athena/api/responses"
	"fmt"
	"net/http"
)

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("called GetUsers")
	user := models.User{}

	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}
