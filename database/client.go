package database

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"examplechat/models"
)

var Instance *gorm.DB
var dbError error

func Connect(connetionString string) () {
	dsn := "host=localhost user=jodeci password=6504 dbname=private_chat port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	Instance, dbError = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("cannot connect to DB")
	}
	log.Println("Connected to DB SUCCESS!")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}