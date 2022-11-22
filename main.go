package main

import (
	"blog-website-backend/database"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env Files")
	} else {
		log.Println("hello from db connect and loading env")
	}

	port := os.Getenv("PORT")
	app := fiber.New()
	app.Listen(":" + port)
}
