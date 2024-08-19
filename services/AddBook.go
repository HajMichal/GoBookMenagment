package services

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"main.go/initializers"
	"main.go/models"
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

	db.Create(book)

	return c.Status(http.StatusOK).JSON(book)
}