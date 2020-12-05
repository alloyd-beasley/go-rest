package main

import "os"

func SetupEnv() {
	os.Setenv("dbHost", DbHost)
	os.Setenv("dbPass", DbPass)
	os.Setenv("dbName", DbName)
	os.Setenv("dbUser", DbUser)
}
