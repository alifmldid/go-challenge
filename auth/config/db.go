package config

import (
	"auth/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbCon *gorm.DB

func InitDB(){
	var err error

	var dbHost = "localhost"
	var dbPort = "5432"
	var dbUser = "developer"
	var dbPass = "devspassword"
	var dbName = "go-blog"

	dbString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)

	// connect to postgres database
	dbCon, err = gorm.Open(postgres.Open(dbString), &gorm.Config{})

	// migrate database
	dbCon.AutoMigrate(&domain.User{})

	if err != nil {
		panic(err.Error())
	}
}

//GetDBConnection
func GetDBConnection() *gorm.DB {
	InitDB()

	return dbCon
}