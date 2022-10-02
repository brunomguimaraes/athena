package controllers

import (
	"athena/api/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "athena2"
)

func (server *Server) Initialize() {
	var err error
	psqlConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	server.DB, err = gorm.Open(postgres.Open(psqlConnectionString), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to database")
		log.Fatal("Error:", err)
	} else {
		fmt.Printf("Connected database")
	}

	server.DB.Debug().AutoMigrate(&models.User{}) //database migration
	// server.DB.Debug().AutoMigrate(&models.User{}, &models.Group{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(address string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(address, server.Router))
}
