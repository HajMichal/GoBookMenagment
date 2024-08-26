package services

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"main.go/initializers"
	"main.go/models"
	"main.go/structs"
	"main.go/utils"
)

func RemoveUser(c fiber.Ctx) error {

	cookieData := new(structs.CookieData)
	if err := c.Bind().Cookie(cookieData); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(err)
	}
	if err := utils.VerifyToken(cookieData.Token); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(err)
	}

	var db = initializers.Db
	id := c.Params("id")

	result := db.Model(&models.User{}).Where("id = ?", id).Delete(&models.User{})
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"error": "User deleted correctly"})
}