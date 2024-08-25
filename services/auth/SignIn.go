package services

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
	"main.go/models"
	services "main.go/services/user"
	"main.go/utils"
)

func SignIn(c fiber.Ctx) error {

	user := new(models.User)

	if err := c.Bind().Body(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "INVALID CREDENTIALS"})
	}

	findUser, err := services.GetUser(c, user.Email)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "WRONG EMAIL"})
	}

	isPwdInvalid := utils.VerifyPwd(user.Password, findUser.Password)
	if isPwdInvalid {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "WRONG PASSWORD"})
	}

	token, err := setToken(c, findUser)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"id": findUser.ID,
		"name": findUser.Name,
		"email": findUser.Email,
		"type": findUser.Type,
		"token": token,
	})
}

func setToken(c fiber.Ctx, user models.User) (string, error) {
	token, err := utils.CreateToken(user.ID, user.Email)
	if err != nil {
		return "", c.Status(http.StatusNotFound).JSON(err.Error())
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.Cookie(cookie)
	return token, nil
}