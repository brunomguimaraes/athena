package controllers

import (
	"athena/api/models"
	"athena/api/responses"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) GetGroups(w http.ResponseWriter, r *http.Request) {
	group := models.Group{}

	groups, err := group.FindAllGroups(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, groups)
}

func (server *Server) GetGroup(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	groupId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	group := models.Group{}
	groupFound, err := group.FindGroupByID(server.DB, uint32(groupId))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, groupFound)
}
