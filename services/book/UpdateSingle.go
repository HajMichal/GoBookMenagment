package services

import (
	"net/http"

	"example.com/m/initializers"
	"example.com/m/models"
	"github.com/gofiber/fiber/v3"
)

func UpdateSingleBook(c fiber.Ctx) error {

	var db = initializers.Db
	book, err := FindBook(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	body := new(models.Book)

	if err := c.Bind().Body(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	if body.ID != 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "You can not change ID"})
	}

	result := db.Model(&book).Updates(body)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(result.Error.Error())
	}

	return c.SendStatus(200)
}