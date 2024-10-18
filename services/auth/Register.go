package services

import (
	"net/http"

	"example.com/m/initializers"
	"example.com/m/models"
	services "example.com/m/services/user"
	"example.com/m/utils"
	"github.com/gofiber/fiber/v3"
)

func Register(c fiber.Ctx) error {
	
	var db = initializers.Db
	user := new(models.User)

	if err := c.Bind().Body(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "INVALID CREDENTIALS"})
	}

	_, err := services.GetUser(c, user.Email)
	if err != nil && err.Error() != "USER NOT FOUND" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "EMAIL IS ALREADY TAKEN"})
	}	

	user.Password, err = utils.HashPwd(user.Password)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	
	result := db.Create(user)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(result.Error.Error())
	}

	// return c.Status(http.StatusCreated).JSON(user)
	return c.SendStatus(http.StatusCreated)
}

