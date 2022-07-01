package main

import (
	"database/sql"
	"fmt"

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
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	// insert
	// // hardcoded
	// insertStmt := `insert into "users"("firstname", "lastname") values('John', 'Wayne')`
	// _, e := db.Exec(insertStmt)
	// CheckError(e)

	// dynamic
	// insertDynStmt := `insert into "users"("id", "firstname", "lastname", "email") values($1, $2, $3, $4)`
	// _, e := db.Exec(insertDynStmt, 3, "Hera", "Guimarães", "hera@dog.com")
	// CheckError(e)

	// Delete
	// deleteStmt := `delete from "users" where id=$1`
	// _, e := db.Exec(deleteStmt, 3)
	// CheckError(e)

	rows, err := db.Query(`SELECT "id", "firstname", "lastname" FROM "users"`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var id int
		var firstname string
		var lastname string

		err = rows.Scan(&id, &firstname, &lastname)
		CheckError(err)

		fmt.Println(id, firstname, lastname)
	}

	CheckError(err)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
