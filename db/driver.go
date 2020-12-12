package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

/*
	DB embeds *sql.DB as Conneciton property.
	Allows us to assign methods to DB struct while also maintaining access to *sql.DB methods.
*/

type DB struct {
	Connection *sql.DB
}

//Initialize database connection.
func Initialize() DB {

	secrets, err := GetSecret()

	if err != nil {
		log.Fatal("Encountered an issue retrieving DB secrets from Secret Manager: ", err.Error())
	}

	host := secrets["host"]
	port := secrets["port"]
	name := secrets["dbname"]
	user := secrets["username"]
	pass := secrets["password"]

	dbConn := fmt.Sprintf("host=%s port=%g user=%s dbname=%s sslmode=require password=%s", host, port, user, name, pass)

	connection, connectionErr := sql.Open("postgres", dbConn)

	if connectionErr != nil {
		log.Fatal("Encountered an issue when establishing initial connection to DB: ", connectionErr.Error())
	}

	db := DB{Connection: connection}

	if pingError := db.Connection.Ping(); pingError != nil {
		log.Fatal("Encountered an issue when testing connection to DB: ", pingError)
	}

	log.Println("Connection to DB was established and tested.")

	return db
}
