package main

import (
	"fmt"
	"go-fintechs-api/database"
	"go-fintechs-api/routes"
)

func main() {
	database.DatabaseConnect()
	fmt.Println("init go server")
	routes.HandleRequest()
}
