package services

import (
	"net/http"

	"example.com/m/initializers"
	"example.com/m/models"
	"example.com/m/structs"
	"example.com/m/utils"
	"github.com/gofiber/fiber/v3"
)

func GetAllBooks(c fiber.Ctx) error {
	var db = initializers.Db
	var books []models.Book

	cookieData := new(structs.CookieData)
	if err := c.Bind().Cookie(cookieData); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(err)
	}
	utils.VerifyToken(cookieData.Token)

	result := db.Find(&books)
	if result.Error != nil {
		c.Status(http.StatusBadRequest).SendString("There was an error during get Books from databse")
	}
	return c.Status(http.StatusOK).JSON(books)
}