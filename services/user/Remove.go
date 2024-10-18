package services

import (
	"net/http"

	"example.com/m/initializers"
	"example.com/m/models"
	"github.com/gofiber/fiber/v3"
)

func RemoveUser(c fiber.Ctx) error {
	var db = initializers.Db
	id := c.Params("id")

	result := db.Model(&models.User{}).Where("id = ?", id).Delete(&models.User{})
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"error": "User deleted correctly"})
}