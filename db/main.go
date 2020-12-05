package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	SetupEnv()
	dbHost := os.Getenv("DbHost")
	dbName := os.Getenv("DbName")
	dbUser := os.Getenv("DbUser")
	dbPass := os.Getenv("DbPass")

	dbConn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=require password=%s", dbHost, 5432, dbUser, dbName, dbPass)
	db, dbErr := sql.Open("postgres", dbConn)
	defer db.Close()

	if dbErr = db.Ping(); dbErr != nil {
		fmt.Println("db ping error")
		fmt.Errorf("%v, ", dbErr)
	}	
}
