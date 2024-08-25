package services

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"main.go/initializers"
	"main.go/models"
)


func UpdateSingleUser(c fiber.Ctx) error {
	
	var db = initializers.Db
	id := c.Params("id")
	userBody := new(models.User)

	if err := c.Bind().Body(userBody); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "INVALID CREDENTIALS"})
	}

	user, err := GetUser(c, id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	result := db.Model(&user).Updates(userBody)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(user)
}