package controllers

import (
	"athena/api/middlewares"
)

func (server *Server) initializeRoutes() {

	//Users routes
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods("POST")
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods("GET")

	//Auth routes
	server.Router.HandleFunc("/auth/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")

	//Market routes
	// server.Router.HandleFunc("/market/items", middlewares.SetMiddlewareJSON(server.GetMarketItems)).Methods("GET")
}
