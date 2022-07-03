package routes

import (
	"athena/api/controllers"
	"athena/api/models"
)

var server = controllers.Server{}

func initializeRoutes() {

	// User Route
	server.Router.HandleFunc("/users/", models.User).Methods("GET")
}
