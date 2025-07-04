package database

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func ConnectToDB(hostname,userName, userPassword, dbName, sslMode string) *gorm.DB {
	//dsn=data source name

	dsn := "host="+ hostname +" user=" + userName + " password=" + userPassword + " dbname=" + dbName + " port=5432 sslmode=" + sslMode +" client_encoding=UTF8"

	var err error
	var DB *gorm.DB
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if(err!=nil){
		log.Fatal("failed to connect to the database",err)
	}
	fmt.Println("Connected to postgres /database using Gorm")
	return DB
}
