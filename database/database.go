package database

import (
	utils "athena/utils"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "athena"
)

func InitializeDB() *gorm.DB {
	utils.PrintMessage("Initializing DB...")

	psqlConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(psqlConnectionString), &gorm.Config{})

	utils.CheckError(err)

	return db

}
