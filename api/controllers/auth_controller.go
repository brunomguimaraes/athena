package controllers

import (
	"net/http"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	// user := models.User{}

	// users, err := user.FindAllUsers(server.DB)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }
	// responses.JSON(w, http.StatusOK, users)
}
