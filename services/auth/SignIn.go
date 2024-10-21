package services

import (
	"net/http"
	"time"

	"example.com/m/models"
	services "example.com/m/services/user"
	"example.com/m/structs"
	"example.com/m/utils"
	"github.com/gofiber/fiber/v3"
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

	isPwdValid := utils.VerifyPwd(user.Password, findUser.Password)
	if !isPwdValid {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "WRONG PASSWORD"})
	}

	token, err := setToken(c, findUser)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(structs.CookieData{
		Id: findUser.ID,
		Name: findUser.Name,
		Email: findUser.Email,
		Type: findUser.Type,
		Token: token,
	})
}

func setToken(c fiber.Ctx, user models.User) (string, error) {
	token, err := utils.CreateToken(user.ID, user.Email, user.Type)
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
