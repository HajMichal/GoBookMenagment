package services

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"main.go/initializers"
	"main.go/models"
)

func DeleteBook(c fiber.Ctx) error {

	var db = initializers.Db
	id := c.Params("id")
	
	result := db.Model(&models.Book{}).Where("id = ?", id).Delete(&models.Book{})
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"success": "Book with id " + " has been deleted"})
}