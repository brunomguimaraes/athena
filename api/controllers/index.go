package controllers

import (
	"athena/api/models"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {
	var err error
	psqlConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

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
