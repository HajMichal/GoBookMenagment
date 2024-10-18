package services

import (
	"net/http"

	"example.com/m/initializers"
	"example.com/m/models"
	"github.com/gofiber/fiber/v3"
)



func AddBook(c fiber.Ctx) error {

	var db = initializers.Db
	book := new(models.Book)

	if err := c.Bind().Body(book); err != nil {
		return err
	}
	
	if book.Title == "" || book.ISBN == "" || book.Author == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	result := db.Create(book)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(book)
}