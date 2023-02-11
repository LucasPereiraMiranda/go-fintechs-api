package main

import (
	"fmt"
	"go-fintechs-api/config"
	"go-fintechs-api/database"
	"go-fintechs-api/routes"
)

func main() {
	host := config.GoDotEnvVariable("DB_HOST")
	user := config.GoDotEnvVariable("DB_USER")
	password := config.GoDotEnvVariable("DB_PASSWORD")
	dbName := config.GoDotEnvVariable("DB_NAME")
	port := config.GoDotEnvVariable("DB_PORT")

	database.DatabaseConnect(host, user, dbName, password, port)
	fmt.Println("init go server")
	routes.HandleRequest()
}
