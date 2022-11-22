package database

import (
	"blog-website-backend/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the .env file")
	}
	dsn := os.Getenv("DSN")

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	} else {
		log.Println("connection successfull")
	}
	DB = database
	database.AutoMigrate(
		&models.User{},
	)

}
