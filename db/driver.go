package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//Initialize database with AWS secret manager
func Initialize() *sql.DB {
	var db *sql.DB
	secrets, err := GetSecret()

	if err != nil {
		fmt.Println("There was a problem retrievng your DB secrets: ", err)
		log.Fatal("Encountered an issue retrieving DB secrets from Secret Manager: ", err)
	}

	host := secrets["host"]
	port := secrets["port"]
	name := secrets["dbname"]
	user := secrets["username"]
	pass := secrets["password"]

	dbConn := fmt.Sprintf("host=%s port=%g user=%s dbname=%s sslmode=require password=%s", host, port, user, name, pass)
	db, dbErr := sql.Open("postgres", dbConn)

	if dbErr != nil {
		log.Fatal("Encountered an issue when establishing initial connection to DB: ", dbErr)
	}

	if pingError := db.Ping(); pingError != nil {
		log.Fatal("Encountered an issue when testing connection to DB: ", pingError)
	} else {
		log.Println("Connection to DB was established and tested.")
	}

	return db
}
