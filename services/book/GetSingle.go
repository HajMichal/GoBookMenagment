package services

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"main.go/initializers"
	"main.go/models"
)

func GetSingleBook(c fiber.Ctx) error {
	
	book, err := FindBook(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(book)
}

func FindBook(c fiber.Ctx) (models.Book, error) {
	var db = initializers.Db
	book := new(models.Book)

	id := c.Params("id")

	result := db.Find(book, id)
	if result.Error != nil {
		return *book, errors.New(result.Error.Error())
	}
	if(book.ID == 0) {
		return *book, errors.New("BOOK WITH THIS ID DOES NOT EXISTS")
	}

	return *book, nil
}