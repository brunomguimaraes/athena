package controllers

import (
	"athena/api/models"
	"athena/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {

	// var err error

	server.DB = database.InitializeDB()

	server.DB.Debug().AutoMigrate(&models.User{}) //database migration

	server.Router = mux.NewRouter()

	// server.initializeRoutes()
}

func (server *Server) Run(address string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(address, server.Router))
}
