package services

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"main.go/initializers"
	"main.go/models"
)

func GetSingleUser(c fiber.Ctx) error {
	id := c.Params("id")
	
	user, err := GetUser(c, id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(user)
}

func GetUser(c fiber.Ctx, credential string) (models.User, error) {
	var db = initializers.Db
	user := new(models.User)

	result := db.First(user, credential)
	if result.Error != nil {
		return *user, errors.New("USER NOT FOUND")
	}
	
	return *user, nil
}

