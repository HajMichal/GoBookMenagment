package utils

import (
	"net/http"

	"example.com/m/structs"
	"github.com/gofiber/fiber/v3"
)

func AdminOnly(c fiber.Ctx) error {
	cookie := new(structs.Cookie)
	if err := c.Bind().Cookie(cookie); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(err)
	}
	userData, err := VerifyToken(cookie.Token)

	if err == nil && userData.Type == 2 {
		return c.Next()
	}
	return c.Status(fiber.StatusForbidden).SendString("AUTHORIZATION FAILED | you have no access to this endpoint ")
}

func EmployeeAuth(c fiber.Ctx) error {
	cookie := new(structs.Cookie)
	if err := c.Bind().Cookie(cookie); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(err)
	}
	userData, err := VerifyToken(cookie.Token)

	if err == nil && userData.Type != 0 {
		return c.Next()
	}
	return c.Status(fiber.StatusForbidden).SendString("AUTHORIZATION FAILED | you have no access to this endpoint ")
}