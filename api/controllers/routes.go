package controllers

import (
	"athena/api/middlewares"
)

func (server *Server) initializeRoutes() {

	//Users routes
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods("GET")
}
