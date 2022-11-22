package controller

import (
	"blog-website-backend/database"
	"blog-website-backend/models"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func validateEmail(email string) bool {
	RE := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return RE.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	// validation
	// check if password is lesser than 6 characters
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be greater than 6 character",
		},
		)
	}
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		return c.JSON(fiber.Map{
			"message": "Invalid Email Address",
		})
	}

	// check if email already exists
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already exist",
		})
	}

	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Phone:     data["phone"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
	}
	user.SetPassword(data["password"].(string))
	err := database.DB.Create(&user)

	if err != nil {
		log.Println()
	}
}
