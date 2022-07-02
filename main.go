package main

import (
	users "athena/api/routes/user"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "athena"
)

func main() {
	router := mux.NewRouter()

	///////////////////////////////
	// Route handles & endpoints //
	//////////////////////////////

	// User Routes
	// Get all users
	router.HandleFunc("/users/", users.GetUsers).Methods("GET")

	// Create a user
	router.HandleFunc("/users/", users.CreateUser).Methods("POST")

	// Delete a specific user by the userID
	// router.HandleFunc("/users/{userid}", users.DeleteUser).Methods("DELETE")

	// Delete all users
	// router.HandleFunc("/users/", users.DeleteAllUsers).Methods("DELETE")

	// serve the app
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
