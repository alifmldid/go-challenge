package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnection() *gorm.DB{
	var dbCon *gorm.DB

	var dbHost = "localhost"
	var dbPort = "5432"
	var dbUser = "developer"
	var dbName = "product"
	var dbPass = "devspassword"

	dbString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)

	// connect to postgres database
	dbCon, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return dbCon
}