package controllers

import (
	"athena/api/middlewares"
)

func (server *Server) initializeRoutes() {

	//Users routes
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods("POST")
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(server.GetUser)).Methods("GET")
	// server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.UpdateUser))).Methods("PUT")
	// server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(server.DeleteUser)).Methods("DELETE")

	//Auth routes
	server.Router.HandleFunc("/auth/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")

	//Market routes
	// server.Router.HandleFunc("/market/items", middlewares.SetMiddlewareJSON(server.GetMarketItems)).Methods("GET")
}
