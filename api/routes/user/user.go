package users

import (
	database "athena/database"
	utils "athena/utils"
	"encoding/json"
	"net/http"
)

// Get all users

type User struct {
	UserID        int    `json:"id"`
	UserFirstName string `json:"firstname"`
	UserLastName  string `json:"lastname"`
	UserEmail     string `json:"email"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []User `json:"data"`
	Message string `json:"message"`
}

// response and request handlers
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := database.InitializeDB()

	utils.PrintMessage("Fetching user data...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM users")

	// check errors
	utils.CheckError(err)

	// var response []JsonResponse
	var users []User

	// Foreach user
	for rows.Next() {
		var id int
		var firstname string
		var lastname string
		var email string

		err = rows.Scan(&id, &firstname, &lastname, &email)

		// check errors
		utils.CheckError(err)

		users = append(users, User{UserID: id, UserFirstName: firstname, UserLastName: lastname, UserEmail: email})
	}

	var response = JsonResponse{Type: "success", Data: users}

	json.NewEncoder(w).Encode(response)
}

// Create a movie

// response and request handlers
func CreateUser(w http.ResponseWriter, r *http.Request) {
	userFirstName := r.FormValue("firstname")
	userLastName := r.FormValue("lastname")
	userEmail := r.FormValue("email")

	var response = JsonResponse{}

	if userFirstName == "" || userLastName == "" || userEmail == "" {
		response = JsonResponse{Type: "error", Message: "User is missing parameter(s)."}
	} else {
		db := database.InitializeDB()

		utils.PrintMessage("Creating user into DB")

		utils.PrintMessage("Inserting new user { firstname: " + userFirstName + " lastname: " + userLastName + " email: " + userEmail + " } ")

		var lastInsertID int
		err := db.QueryRow("INSERT INTO users(userFirstName, userLastName, userEmail) VALUES($1, $2, $3, $4) returning id;", userFirstName, userLastName, userEmail).Scan(&lastInsertID)

		// check errors
		utils.CheckError(err)

		response = JsonResponse{Type: "success", Message: "User has been created successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
