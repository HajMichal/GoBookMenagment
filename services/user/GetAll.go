package services

import (
	"net/http"

	"example.com/m/initializers"
	"example.com/m/models"
	"github.com/gofiber/fiber/v3"
)

func GetAllUsers(c fiber.Ctx) error {

	var db = initializers.Db
	var users []models.User

	result := db.Find(&users)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(users)
}