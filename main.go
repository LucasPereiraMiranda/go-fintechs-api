package main

import (
	"fmt"
	"go-fintechs-api/database"
	"go-fintechs-api/env"
	"go-fintechs-api/routes"
)

func main() {
	host := env.GoDotEnvVariable("DB_HOST")
	user := env.GoDotEnvVariable("DB_USER")
	password := env.GoDotEnvVariable("DB_PASSWORD")
	dbName := env.GoDotEnvVariable("DB_NAME")
	port := env.GoDotEnvVariable("DB_PORT")
	database.DatabaseConnect(host, user, dbName, password, port)
	fmt.Println("init go server")
	routes.HandleRequest()
}
