package database

import (
	utils "athena/utils"
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "athena"
)

func InitializeDB() *sql.DB {
	utils.PrintMessage("Initializing DB...")
	psqlConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlConnectionString)
	utils.CheckError(err)

	return db

}
