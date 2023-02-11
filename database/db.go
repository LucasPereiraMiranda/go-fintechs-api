package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnect(host string, user string, databaseName string, password string, port string) {

	connectString := "host= " + host + " user=" + user + " password=" + password + " dbname=" + databaseName + " port=" + port + " sslmode=disable"

	DB, err = gorm.Open(postgres.Open(connectString))
	if err != nil {
		log.Panic("database connection error")
	}
}
